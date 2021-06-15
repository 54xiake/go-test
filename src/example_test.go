package libfoo

import (
	"testing"
)

func add(x, y int) int {
	return x + y
}

func TestFoo(t *testing.T) {
	if add(1, 2) != 3 {
		t.Error("test foo:Addr failed")
	} else {
		t.Log("test foo:Addr pass")
	}

	//i := Foo(1)
	//fmt.Printf("%v\n",i)
}
