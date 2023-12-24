package performance

import (
	"fmt"
	"testing"
)

func TestAbc(t *testing.T) {
	t.Log("123")
}

// go test -bench=. -run=none -benchtime=3s
// go test -bench=. -benchmem -run=none 来查看每次操作分配内存的次数
func BenchmarkAbc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprint("Parallel")
	}
}

// go test -bench=. -benchmem -run=BenchmarkSprints
func BenchmarkSprints(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// do something
			fmt.Sprint("Parallel")
		}
	})
}
