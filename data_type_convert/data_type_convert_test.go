package data_type_convert

import (
	"fmt"
	"testing"
)

func Test_StringToInt(t *testing.T) {
	s := "123"
	t.Log(StringToInt(s))
}

func Test_PrintSpeicalString(t *testing.T) {
	cs1 := "abc\n123"
	t.Log(PrintSpeicalString(cs1))
}

func Test_IntToT(t *testing.T) {
	x := BaseConvert(5, 2)
	fmt.Println(x)
}
