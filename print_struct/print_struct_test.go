package print_struct

import "testing"

type People struct {
	Name string
	Age  int
}

func NewPeople(name string, age int) *People {
	return &People{
		Name: name,
		Age:  age,
	}
}

var p1 = NewPeople("Jone", 13)

func Test_PrintStructString(t *testing.T) {
	p1_str := PrintStructString(p1)
	t.Logf("PrintStructString==>[\n%s\n] \n ", string(p1_str))
}

func Test_PrintStructJson(t *testing.T) {
	p1_str := PrintStructJson(p1)
	t.Logf("PrintStructJson==>[\n%s\n] \n ", string(p1_str))
}
