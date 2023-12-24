package main

/*
int foo(int c){
   return c+1;
};

extern int foo(int a);

#cgo CFLAGS: -I ./
#cgo LDFLAGS: -L ${SRCDIR}/ -l foo
#include <foo.h>
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.foo(2))
}
