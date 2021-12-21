package data_type_convert

import "testing"

func Test_StringToInt(t *testing.T) {
	s := "123"
	t.Log(StringToInt(s))
}
