package time_convert

import (
	"fmt"
	"time"
)

/*
time to timestemp
*/
func TimeConvTimestemp(format, acturalTime string) int64 {
	startTime, err := time.Parse(format, acturalTime)
	if err != nil {
		fmt.Printf("TimeConvTimestemp==>time.Parse==>failed==>[%s] \n", err.Error())
	}
	timestemp := startTime.Unix()
	// fmt.Printf("TimeConvTimestemp==>[%d] \n", timestemp)
	return timestemp
}
