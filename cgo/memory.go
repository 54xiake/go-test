package main

/*
//#include <stdlib.h>
//#include <stdio.h>

extern char* NewGoString(char* );
extern void FreeGoString(char* );
extern void PrintGoString(char* );

//void* makeslice(size_t memsize) {
//	return malloc(memsize);
//}

static void printString(char* s) {
	char* gs = NewGoString(s);
	PrintGoString(gs);
	FreeGoString(gs);
}
*/
import "C"
import "unsafe"
import "sync"

type ObjectId int32

var refs struct {
	sync.Mutex
	objs map[ObjectId]interface{}
	next ObjectId
}

func init() {
	refs.Lock()
	defer refs.Unlock()

	refs.objs = make(map[ObjectId]interface{})
	refs.next = 1000
}

func NewObjectId(obj interface{}) ObjectId {
	refs.Lock()
	defer refs.Unlock()

	id := refs.next
	refs.next++

	refs.objs[id] = obj
	return id
}

func (id ObjectId) IsNil() bool {
	return id == 0
}

func (id ObjectId) Get() interface{} {
	refs.Lock()
	defer refs.Unlock()

	return refs.objs[id]
}

func (id *ObjectId) Free() interface{} {
	refs.Lock()
	defer refs.Unlock()

	obj := refs.objs[*id]
	delete(refs.objs, *id)
	*id = 0

	return obj
}

//func makeByteSlize(n int) []byte {
//	p := C.makeslice(C.size_t(n))
//	return ((*[1 << 31]byte)(p))[0:n:n]
//}
//
//func freeByteSlice(p []byte) {
//	C.free(unsafe.Pointer(&p[0]))
//}

//export NewGoString
func NewGoString(s *C.char) *C.char {
	gs := C.GoString(s)
	id := NewObjectId(gs)
	return (*C.char)(unsafe.Pointer(uintptr(id)))
}

//export FreeGoString
func FreeGoString(p *C.char) {
	id := ObjectId(uintptr(unsafe.Pointer(p)))
	id.Free()
}

//export PrintGoString
func PrintGoString(s *C.char) {
	id := ObjectId(uintptr(unsafe.Pointer(s)))
	gs := id.Get().(string)
	print(gs)
}

func main() {
	//n := 1 << 31
	//s := makeByteSlize(n)
	//s[len(s)-1] = 1
	//print(s[len(s)-1])
	//freeByteSlice(s)

	str := "hello"
	//C.printString((*C.char)(unsafe.Pointer(&str)))
	C.printString(C.CString(str))
	//C.printString("hello")

}
