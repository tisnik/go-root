// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam příkladů z osmdesáté páté části:
//    https://github.com/tisnik/go-root/blob/master/article_85/README.md

package main

import "C"

import "fmt"

//export hello
func hello(name *C.char) int {
	goName := C.GoString(name)
	fmt.Printf("Hello %s\n", goName)
	return len(goName)
}

func main() {}
