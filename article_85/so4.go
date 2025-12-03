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

//export add
func add(x, y int64) int64 {
	result := x + y
	fmt.Printf("Called add(%d, %d) with result %d\n", x, y, result)
	return result
}

func main() {}
