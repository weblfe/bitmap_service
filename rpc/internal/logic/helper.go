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

func GetUserDayKeepByBitmapArr(userBitmapArr []*roaring64.Bitmap) (keep float32, total int, left int) {
	var (
		n uint64 = 0
		l uint64 = 0
	)
	if userBitmapArr == nil {
		return keep, int(n), int(l)
	}
	// bitmap
	if len(userBitmapArr) >= 2 {
		//fmt.Println(fmt.Sprintf("first-day:%d", count64(userBitmapArr[0].ToArray())))
		n = userBitmapArr[1].GetCardinality()
		// fmt.Println(fmt.Sprintf("first-day:%d", userBitmapArr[1].GetCardinality()))
		if n != 0 {
			// 交集合
			userBitmapArr[0].And(userBitmapArr[1])
			if userBitmapArr[0] != nil {
				l = userBitmapArr[0].GetCardinality()
				// fmt.Println(fmt.Sprintf("last-day:%d", count64(userBitmapArr[0].ToArray())))
				// fmt.Println(fmt.Sprintf("last-day:%d", l))
				keep = float32(l) / float32(n) * 100
				// fmt.Println(fmt.Sprintf("keep:%.2f", keep))
			}
		}
	}
	return keep, int(n), int(l)
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
	case 360:
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
