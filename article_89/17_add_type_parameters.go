// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z osmdesáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_89/README.md

package main

import "fmt"

type numeric interface {
	int | float64
}

func add[T numeric](x T, y T) T {
	return x + y
}

func main() {
	fmt.Println(add(1, 3.14))
	fmt.Println(add(3.14, 1))
}
