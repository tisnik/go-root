// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z osmdesáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_89/README.md

package main

import "fmt"

type Slice[T any] []T

func (s Slice[T]) Print() {
	for _, value := range s {
		fmt.Println(value)
	}
}

func Print[T any](s Slice[T]) {
	for _, value := range s {
		fmt.Println(value)
	}
}

func main() {
	s := Slice[int]{1, 2, 3}

	Print(s)

	fmt.Println()

	s.Print()
}
