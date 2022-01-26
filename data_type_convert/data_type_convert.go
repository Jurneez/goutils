package data_type_convert

import (
	"strconv"
)

// string to int
func StringToInt(str string) int {
	int, _ := strconv.Atoi(str)
	return int
}

func PrintSpeicalString(str string) string {
	return strconv.Quote(str)
}

// 10 base to t base
func BaseConvert(number, base int) int {
	if number == 0 {
		return 0
	}
	if number < 0 {
		return -1 // 目前不支持负数
	}
	remainders := make([]int, 0)
	for {
		if number < base {
			remainders = append(remainders, number)
			break
		}
		remainders = append(remainders, number%base)
		number = number / base
	}
	count := len(remainders)
	result := 0
	for i := 0; i < count; i++ {
		k := count - 1 - i
		result = result + remainders[k]*power(10, k)
	}
	return result
}

func power(number, times int) int {
	result := 1
	for i := 0; i < times; i++ {
		result = result * number
	}
	return result
}
