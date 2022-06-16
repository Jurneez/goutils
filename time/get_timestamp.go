package time

import (
	"time"
)

// 获取当天开始时间戳
func GetStartTimeStampOfDay() int64 {
	y, m, d := time.Now().Date()
	startTime := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	return startTime.Unix()
}

// 获取当天结束时间戳
func GetEndTimeStampOfDay() int64 {
	y, m, d := time.Now().Date()
	endTime := time.Date(y, m, d, 23, 59, 59, 0, time.Local)
	return endTime.Unix()
}

// 获取指定年-月月初时间戳
func GetMonthStartTimeStamp(year int, month time.Month) int64 {
	// 指定年-月第一天
	specifyMonthStart := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	return specifyMonthStart.Unix()
}

// 获取指定年-月 月末时间戳
func GetMonthEndTimeStamp(year int, month time.Month) int64 {
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	// 下个月减1为本月的最后一天
	end := thisMonth.AddDate(0, 1, -1)
	thisMonthEnd := time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 0, time.Local)
	return thisMonthEnd.Unix()
}
