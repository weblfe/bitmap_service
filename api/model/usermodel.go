package model

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/syncx"
	"github.com/weblfe/bitmap_service/api/internal/config"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	exclusiveCalls                         = syncx.NewSharedCalls()
	stats                                  = cache.NewStat("redis")
	userDayLoginFieldNames                 = builderx.RawFieldNames(&UserDayLogin{})
	userDayLoginRows                       = strings.Join(userDayLoginFieldNames, ",")
	userDayLoginRowsExpectAutoSet          = strings.Join(stringx.Remove(userDayLoginFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userDayLoginRowsWithPlaceHolder        = strings.Join(stringx.Remove(userDayLoginFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
	cacheUserDayLoginIdPrefix              = "cache#userDayLogin#id#"
	cacheUserDayLoginDatePrefix            = "cache#userDayLogin#date#"
	cacheBitmapUserDayLoginDatePrefix      = "cache#bitmap#userDayLogin#date#"
	cacheStartBitmapUserDayLoginDatePrefix = "cache#bitmap#start#userDayLogin#date#"
)

const (
	TimeFormat               = "2006-01-02T15:04:05Z07:00"
	DateFormat               = "2006-01-02"
	DefaultSafeTime          = time.Minute * 3
	RoleAnchor               = "1" // 主播
	RoleNormal               = "0" // 普通用户
	DateEnd         DateType = 0   // 截止时间类型
	DateBegin       DateType = 1   // 开始时间类型
)

type (
	DateType          int
	UserDayLoginModel interface {
		Insert(data UserDayLogin) (sql.Result, error)
		FindOne(id int64) (*UserDayLogin, error)
		Update(data UserDayLogin) error
		Delete(id int64) error
		GetUserArr(queryMap QueryMap) ([][]uint64, error)
		GetUserBitMapArr(queryMap QueryMap) ([]*roaring64.Bitmap, error)
	}

	defaultUserDayLoginModel struct {
		CachedConn
		table  string
		config config.Config
	}

	UserDayLogin struct {
		Id              int64     `db:"id"`                // ID
		Date            time.Time `db:"date"`              // 日期
		Uid             int64     `db:"uid"`               // Uid
		Count           int64     `db:"count"`             // 登陆次数
		IsAnchor        int64     `db:"is_anchor"`         // 是否认证主播(0:非主播,1:主播)
		IsNewer         int64     `db:"is_newer"`          //  是否新用户
		UserClass       int64     `db:"user_class"`        // 用户类型(0:非付费用户,1:付费用户)
		StartActivityAt time.Time `db:"start_activity_at"` // 当天首次活跃时间点
		LastActivityAt  time.Time `db:"last_activity_at"`  // 当天最后一次活跃时间点
		CreatedAt       time.Time `db:"created_at"`        // 创建时间
		UpdatedAt       time.Time `db:"updated_at"`        // 更新时间
	}

	QueryMap struct {
		Day     int      `json:"day"`
		Channel string   `json:"channel,omitempty"`
		Role    string   `json:"role,omitempty"`
		Date    string   `json:"data,omitempty"`
		Type    DateType `json:"type,omitempty"`
	}

	uidResult struct {
		Uid int64 `db:"uid"` // Uid
	}

	countResult struct {
		Count int64 `db:"count"` // Count
	}

	CacheResultUser struct {
		Uids      []uint64  `json:"uids"`
		ExpiredAt time.Time `json:"expired_at"`
	}

	CacheBitmapResultUser struct {
		UidBitMap []byte    `json:"uid_bit_map"`
		ExpiredAt time.Time `json:"expired_at"`
		Count     uint64    `json:"count"`
	}

	CachedConn struct {
		sqlc.CachedConn
		cacheProvider cache.Cache
	}
)

func NewUserDayLoginModel(conn sqlx.SqlConn, table string, c cache.CacheConf, cnf config.Config) UserDayLoginModel {
	if cnf.SafeTime == 0 {
		cnf.SafeTime = DefaultSafeTime
	}
	return &defaultUserDayLoginModel{
		CachedConn: CachedConn{
			CachedConn:    sqlc.NewConn(conn, c),
			cacheProvider: NewCacheProvider(c),
		},
		table:  fmt.Sprintf("`%s`", table),
		config: cnf,
	}
}

func (m *defaultUserDayLoginModel) Insert(data UserDayLogin) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userDayLoginRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Date, data.Uid, data.Count, data.IsAnchor, data.UserClass, data.StartActivityAt, data.LastActivityAt, data.CreatedAt, data.UpdatedAt)

	return ret, err
}

func (m *defaultUserDayLoginModel) FindOne(id int64) (*UserDayLogin, error) {
	UserDayLoginIdKey := fmt.Sprintf("%s%v", cacheUserDayLoginIdPrefix, id)
	var resp UserDayLogin
	err := m.QueryRow(&resp, UserDayLoginIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userDayLoginRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserDayLoginModel) Update(data UserDayLogin) error {
	UserDayLoginIdKey := fmt.Sprintf("%s%v", cacheUserDayLoginIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userDayLoginRowsWithPlaceHolder)
		return conn.Exec(query, data.Date, data.Uid, data.Count, data.IsAnchor, data.UserClass, data.StartActivityAt, data.LastActivityAt, data.CreatedAt, data.UpdatedAt, data.Id)
	}, UserDayLoginIdKey)
	return err
}

func (m *defaultUserDayLoginModel) Delete(id int64) error {
	UserDayLoginIdKey := fmt.Sprintf("%s%v", cacheUserDayLoginIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, UserDayLoginIdKey)
	return err
}

func (m *defaultUserDayLoginModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserDayLoginIdPrefix, primary)
}

func (m *defaultUserDayLoginModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userDayLoginRows, m.table)
	return conn.QueryRow(v, query, primary)
}

func (m *defaultUserDayLoginModel) GetUserArr(where QueryMap) ([][]uint64, error) {
	var date, err = time.Parse(DateFormat, where.Date)
	if assertError(err) {
		return nil, err
	}

	var (
		arr        [][]uint64
		isCached   bool
		one        = make([]uidResult, 10)
		result     []uidResult
		counter    = &countResult{}
		cacheUser  = &CacheResultUser{}
		day        = where.Day
		channel    = where.Channel
		role       = where.Role
		countQuery = ""
		userQuery  = ""
		nextQuery  = ""
	)
	dateStr := date.Format(DateFormat)
	// 查询当天登录用户
	userQuery = fmt.Sprintf("select %s from %s where date='%s' ", "uid", m.table, dateStr)
	// 查询渠道当天登录用户
	if channel != "" {
		userQuery = fmt.Sprintf("%s and reg_channel=%s", userQuery, channel)
	}
	// 身份类型
	switch role {
	case RoleAnchor:
		// 主播
		userQuery = fmt.Sprintf("%s and is_anchor=1 ", userQuery)
	case RoleNormal:
		// 普通用户
		userQuery = fmt.Sprintf("%s and is_anchor=0 ", userQuery)
	default:
		// 全部
		userQuery = fmt.Sprintf("%s ", userQuery)
	}
	// 用户数量查询
	countQuery = fmt.Sprintf("select count(`uid`) as `count` from (%s) `t`", userQuery)
	if err = m.QueryRowNoCache(counter, countQuery); err != nil {
		return nil, err
	}
	key := fmt.Sprintf("%v%s%d", cacheUserDayLoginDatePrefix+role+channel, dateStr, counter.Count)
	arr = append(arr, []uint64{})
	err = m.CachedConn.GetCache(key, cacheUser)
	if err == nil && !cacheUser.ExpiredAt.Before(time.Now()) {
		isCached = true
		arr[0] = cacheUser.Uids
	} else {
		if err = m.QueryRowsNoCache(&one, userQuery); err != nil {
			return nil, err
		}
		arr[0] = append(arr[0], toArray64(one)...)
	}

	if !isCached {
		err = m.CachedConn.SetCache(key, &CacheResultUser{
			ExpiredAt: time.Now().Add(m.config.SafeTime),
			Uids:      arr[0],
		})
		assertError(err)
	}

	d := time.Hour * 24 * time.Duration(day)
	date = date.Add(-d)
	nextQuery = fmt.Sprintf("select %s from %s where date='%s'", "uid", m.table, date.Format(DateFormat))

	switch role {
	case RoleAnchor:
		nextQuery = fmt.Sprintf("%s and is_anchor=1", nextQuery)
	case RoleNormal:
		nextQuery = fmt.Sprintf("%s and is_anchor=0 and is_newer=1", nextQuery)
	default:
		nextQuery = fmt.Sprintf("%s and is_newer=1", nextQuery)
	}

	err = m.QueryRowsNoCache(&result, nextQuery)
	arr = append(arr, []uint64{})
	if err == nil {
		arr[1] = append(arr[1], toArray64(result)...)
	}
	return arr, nil
}

func (m *defaultUserDayLoginModel) GetUserBitMapArr(where QueryMap) ([]*roaring64.Bitmap, error) {
	var date, err = time.Parse(DateFormat, where.Date)
	if assertError(err) {
		return nil, err
	}
	var (
		day      = where.Day
		role     = where.Role
		dateType = where.Type
		channel  = where.Channel
		arr      []*roaring64.Bitmap
	)
	// 当天登录数据
	dateStr := date.Format(DateFormat)
	switch dateType {
	case DateBegin:
		arr, err = m.getByStartDateType(dateStr, role, channel, day, date)
	case DateEnd:
		arr, err = m.getByEndDateType(dateStr, role, channel, day, date)
	}
	return arr, nil
}

// 获取用户登录 bitmap
func (m *defaultUserDayLoginModel) getUserBitMap(date string, counter *countResult, prefix string, query string) (*roaring64.Bitmap, error) {
	var (
		err       error
		isCached  bool
		result    []uidResult
		bitmap    *roaring64.Bitmap
		cacheUser = &CacheBitmapResultUser{}
	)
	key := fmt.Sprintf("%v%s", prefix, date)
	err = m.CachedConn.GetCache(key, cacheUser)
	if err == nil && !cacheUser.ExpiredAt.Before(time.Now()) && int64(cacheUser.Count) == counter.Count {
		isCached = true
		bitmap = roaring64.New()
		if n, err := bitmap.ReadFrom(bytes.NewBuffer(cacheUser.UidBitMap)); assertError(err) || n < 0 {
			return nil, fmt.Errorf("%s", "bitmap readFrom error")
		}

	} else {
		if err = m.QueryRowsNoCache(&result, query); assertError(err) {
			return nil, err
		}
		logx.Infof(query, " --> ", result)
		bitmap = roaring64.BitmapOf(toArray64(result)...)
	}
	bitmap.RunOptimize()
	if !isCached && !bitmap.IsEmpty() {
		expiredAt := time.Now().Add(m.config.SafeTime)
		err = m.CachedConn.GetCacheProvider().SetWithExpire(key, &CacheBitmapResultUser{
			ExpiredAt: expiredAt,
			UidBitMap: toBytes(bitmap),
			Count:     uint64(counter.Count),
		}, m.config.SafeTime)
		assertError(err)
	}
	return bitmap, nil
}

func (m *defaultUserDayLoginModel) ClearCache() bool {
	return false
}

// 截止日期往前查询
func (m *defaultUserDayLoginModel) getByEndDateType(dateStr string, role string, channel string, day int, date time.Time) ([]*roaring64.Bitmap, error) {

	var (
		err              error
		countQuery       = ""
		userQuery        = ""
		nextCountQuery   = ""
		nextUserQuery    = ""
		counter          = &countResult{}
		arr              = make([]*roaring64.Bitmap, 0, 2)
		cachePrefix      = m.GetCacheKey(cacheBitmapUserDayLoginDatePrefix, role, channel)
		cacheStartPrefix = m.GetCacheKey(cacheStartBitmapUserDayLoginDatePrefix, role, channel)
	)

	userQuery = fmt.Sprintf("select %s from %s where date='%s'", "uid", m.table, dateStr)
	if channel != "" {
		userQuery = fmt.Sprintf("%s and reg_channel='%s'", userQuery, channel)
	}
	// 身份类型
	switch role {
	case RoleAnchor:
		// 主播
		userQuery = fmt.Sprintf("%s and is_anchor=1 ", userQuery)
	case RoleNormal:
		// 普通用户
		userQuery = fmt.Sprintf("%s and is_anchor=0 ", userQuery)
	default:
		// 全部
		userQuery = fmt.Sprintf("%s", userQuery)
	}
	// 数量变化
	countQuery = fmt.Sprintf("select count(`uid`) as `count` from (%s) `t`", userQuery)
	if err = m.QueryRowNoCache(counter, countQuery); err != nil {
		return nil, err
	}
	logx.Infof(userQuery, " -- ", counter.Count)
	// 无
	if counter.Count < 0 {
		return nil, nil
	}
	arr[0], err = m.getUserBitMap(dateStr, counter, cachePrefix, userQuery)
	if assertError(err) {
		return arr, err
	}
	// 查询用户注册情况(1-30)天
	counter.Count = 0
	// 往前推
	date = date.Add(-time.Hour * 24 * time.Duration(day))
	nextDate := date.Format(DateFormat)
	nextUserQuery = fmt.Sprintf("select %s from %s where date='%s' ", "uid", m.table, nextDate)
	if channel != "" {
		nextUserQuery = fmt.Sprintf("%s and reg_channel='%s'", nextUserQuery, channel)
	}
	// 主播留存
	switch role {
	case RoleAnchor:
		// 主播
		nextUserQuery = fmt.Sprintf("%s and is_anchor=1 ", nextUserQuery)
		nextCountQuery = fmt.Sprintf("select count(`uid`) as `count` from (%s) `t`", nextUserQuery)
	case RoleNormal:
		// 普通用户
		nextUserQuery = fmt.Sprintf("%s and is_anchor=0 and is_newer=1", nextUserQuery)
		nextCountQuery = fmt.Sprintf("select count(`uid`) as `count` from (%s) `t`", nextUserQuery)
	default:
		// 全部
		nextUserQuery = fmt.Sprintf("%s and is_newer=1 ", nextUserQuery)
		nextCountQuery = fmt.Sprintf("select count(`uid`) as `count` from (%s) `t`", nextUserQuery)
	}
	if err = m.QueryRowNoCache(counter, nextCountQuery); err != nil {
		return arr, err
	}
	arr[1], err = m.getUserBitMap(nextDate, counter, cacheStartPrefix, nextUserQuery)
	if assertError(err) {
		return arr, err
	}
	return arr, nil
}

// GetCacheKey 获取缓存key
func (m *defaultUserDayLoginModel) GetCacheKey(args ...string) string {
	var values []string
	for _, v := range args {
		if v == "" {
			continue
		}
		values = append(values, v)
	}
	return strings.Join(values, ":")
}

// 开始日期往后查询
func (m *defaultUserDayLoginModel) getByStartDateType(dateStr string, role string, channel string, day int, date time.Time) ([]*roaring64.Bitmap, error) {
	var (
		err              error
		countQuery       = ""
		userQuery        = ""
		nextCountQuery   = ""
		nextUserQuery    = ""
		counter          = &countResult{}
		arr              = make([]*roaring64.Bitmap, 2)
		cachePrefix      = m.GetCacheKey(cacheBitmapUserDayLoginDatePrefix, role, channel)
		cacheStartPrefix = m.GetCacheKey(cacheStartBitmapUserDayLoginDatePrefix, role, channel)
	)
	userQuery = fmt.Sprintf("select %s from %s where date='%s'", "uid", m.table, dateStr)
	if channel != "" {
		userQuery = fmt.Sprintf("%s and reg_channel='%s'", userQuery, channel)
	}
	// 身份类型
	switch role {
	case RoleAnchor:
		// 主播
		userQuery = fmt.Sprintf("%s and is_anchor=1 ", userQuery)
	case RoleNormal:
		// 普通用户
		userQuery = fmt.Sprintf("%s and is_anchor=0 and is_newer=1 ", userQuery)
	default:
		// 全部
		userQuery = fmt.Sprintf("%s and is_newer=1 ", userQuery)
	}
	// 数量变化
	countQuery = fmt.Sprintf("select count(`uid`) as `count` from (%s) `t`", userQuery)
	if err = m.QueryRowNoCache(counter, countQuery); err != nil {
		return nil, err
	}
	logx.Infof(userQuery, " -- ", counter.Count)
	// 无
	if counter.Count < 0 {
		return nil, nil
	}

	arr[0], err = m.getUserBitMap(dateStr, counter, cacheStartPrefix, userQuery)
	if assertError(err) {
		return arr, err
	}
	// 查询用户注册情况(1-30)天
	counter.Count = 0
	// 往后退推
	date = date.Add(time.Hour * 24 * time.Duration(day))
	nextDate := date.Format(DateFormat)
	nextUserQuery = fmt.Sprintf("select %s from %s where date='%s' ", "uid", m.table, nextDate)
	if channel != "" {
		nextUserQuery = fmt.Sprintf("%s and reg_channel='%s'", nextUserQuery, channel)
	}
	// 主播留存
	switch role {
	case RoleAnchor:
		// 主播
		nextUserQuery = fmt.Sprintf("%s and is_anchor=1 ", nextUserQuery)
		nextCountQuery = fmt.Sprintf("select count(`uid`) as `count` from (%s) `t`", nextUserQuery)
	case RoleNormal:
		// 普通用户
		nextUserQuery = fmt.Sprintf("%s and is_anchor=0", nextUserQuery)
		nextCountQuery = fmt.Sprintf("select count(`uid`) as `count` from (%s) `t`", nextUserQuery)
	default:
		// 全部
		nextUserQuery = fmt.Sprintf("%s", nextUserQuery)
		nextCountQuery = fmt.Sprintf("select count(`uid`) as `count` from (%s) `t`", nextUserQuery)
	}
	if err = m.QueryRowNoCache(counter, nextCountQuery); err != nil {
		return arr, err
	}
	arr[1], err = m.getUserBitMap(nextDate, counter, cachePrefix, nextUserQuery)
	if assertError(err) {
		return arr, err
	}
	return arr, nil
}

// 数组
func toArray(arr []uidResult) []uint32 {
	var intArr []uint32
	for _, it := range arr {
		intArr = append(intArr, uint32(it.Uid))
	}
	return intArr
}

// 数组
func toArray64(arr []uidResult) []uint64 {
	var intArr []uint64
	for _, it := range arr {
		intArr = append(intArr, uint64(it.Uid))
	}
	return intArr
}

func assertError(err error) bool {
	if err != nil {
		logx.Error(err)
		return true
	}
	return false
}

func toString(bitmap *roaring64.Bitmap) string {
	if bitmap == nil {
		return ""
	}
	str, err := bitmap.ToBytes()
	if assertError(err) {
		return ""
	}
	return string(str)
}

func toBytes(bitmap *roaring64.Bitmap) []byte {
	if bitmap == nil {
		return nil
	}
	byteArr, err := bitmap.ToBytes()
	if assertError(err) {
		return nil
	}
	return byteArr
}

func NewCacheProvider(c cache.CacheConf, opts ...cache.Option) cache.Cache {
	return cache.New(c, exclusiveCalls, stats, sql.ErrNoRows, opts...)
}

func (c *CachedConn) GetCacheProvider() cache.Cache {
	return c.cacheProvider
}

func ParserDateType(ty int) DateType {
	var d = DateType(ty)
	if d != DateBegin && d != DateEnd {
		return DateEnd
	}
	return d
}
