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
func hello(name string) int {
	fmt.Printf("Hello %s\n", name)
	return len(name)
}

func main() {}
