// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam příkladů z osmdesáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_86/README.md

package main

// #include <stdlib.h>
import "C"
import "unsafe"

//export concat
func concat(text1, text2 *C.char) *C.char {
	t1 := C.GoString(text1)
	t2 := C.GoString(text2)

	result := t1 + t2
	return C.CString(result)
}

//export freeString
func freeString(s *C.char) {
	C.free(unsafe.Pointer(s))
}

func main() {}
