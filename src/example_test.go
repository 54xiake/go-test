package libfoo

import (
	"fmt"
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

// go test -bench=. -run=BenchmarkFoo
func BenchmarkFoo(b *testing.B) {
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		if add(1, 2) != 3 {
			fmt.Println("err")
		} else {
			fmt.Println("ok")
		}
	}
}
