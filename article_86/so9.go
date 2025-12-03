// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam příkladů z osmdesáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_86/README.md

package main

import "C"
import "unsafe"

//export sum
func sum(values *C.int, length int) int64 {
	var sum int64 = 0
	slice := unsafe.Slice(values, length)
	for _, value := range slice {
		sum += int64(value)
	}
	return sum
}

func main() {}
