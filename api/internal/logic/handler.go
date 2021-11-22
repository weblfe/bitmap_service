package logic

import (
	"fmt"
	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/weblfe/bitmap_service/api/model"
	"time"
)

const (
	dayName       = "日"
	secondDayName = "次"
	weekName      = "周"
	monthName     = "30天"
	yearName      = "年"
	keeperName    = "留存"
)

func GetUserDayKeep(userArr [][]uint64) (float32, int, int) {
	var (
		bits          = make([][]uint64, 2)
		keepBitMapArr []*roaring64.Bitmap
	)
	for i, users := range userArr {
		bits[i] = make([]uint64, cap(users))
		for _, id := range users {
			bits[i] = append(bits[i], uint64(id))
		}
		it := roaring64.BitmapOf(bits[i]...)
		it.RunOptimize()
		keepBitMapArr = append(keepBitMapArr, it)
	}
	// bitmap
	return GetUserDayKeepByBitmapArr(keepBitMapArr)
}

// GetUserDayKeepByBitmapArr 获取留存数据
// keep :留存率 , total 对应天总注册人数, left 对应天注册留存人数
func GetUserDayKeepByBitmapArr(userBitmapArr []*roaring64.Bitmap) (keep float32, total int, left int) {
	var (
		totalNum uint64 = 0
		leftNum  uint64 = 0
	)
	if userBitmapArr == nil {
		return keep, 0, 0
	}
	// bitmap
	if len(userBitmapArr) >= 2 {
		totalNum = userBitmapArr[1].GetCardinality()
		if totalNum != 0 && userBitmapArr[0] != nil {
			// 1001100 & 1001001
			userBitmapArr[0].And(userBitmapArr[1])
			leftNum = userBitmapArr[0].GetCardinality()
			keep = float32(leftNum) / float32(totalNum) * 100
		}
	}
	return keep, int(totalNum), int(leftNum)
}

func getName(date string, day int64, channel string) string {
	var d, err = time.Parse(model.DateFormat, date)
	if err != nil {
		return ""
	}
	if channel == "" {
		return fmt.Sprintf("%s (%s) "+keeperName, transformDay(day), d.Add(-time.Hour*24*time.Duration(day)).Format(model.DateFormat))
	}
	return fmt.Sprintf("%s-%s (%s) "+keeperName, channel, transformDay(day), d.Add(-time.Hour*24*time.Duration(day)).Format(model.DateFormat))
}

func transformDay(day int64) string {
	// %d天
	switch day {
	case 1:
		return secondDayName
	case 7:
		return weekName
	case 30:
		return monthName
	case 360,365:
		return yearName
	}
	return fmt.Sprintf("%d%s", day, dayName)
}

func count(arr []uint32) int {
	var n = 0
	for _, v := range arr {
		if v > 0 {
			n++
		}
	}
	return n
}

func count64(arr []uint64) int {
	var n = 0
	for _, v := range arr {
		if v > 0 {
			n++
		}
	}
	return n
}
