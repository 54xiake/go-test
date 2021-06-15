// +build debug
package main

//#include <stdio.h>
import "C"
import (
	cgo "github.com/54xiake/gotest/cgo-c++"
	"unsafe"
)

var buildMode = "debug"

func main() {
	buf := cgo.NewMyBuffer(1024)
	defer buf.Delete()

	copy(buf.Data(), []byte("hello\x00"))
	C.puts((*C.char)(unsafe.Pointer(&(buf.Data()[0]))))
}
