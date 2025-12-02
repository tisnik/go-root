// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z devadesáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_96/README.md

package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

func main() {
	var q deque.Deque[int]

	fmt.Println("Length Capacity")
	for i := 1; i <= 1000; i++ {
		q.PushBack(i)
		fmt.Printf("%5d %5d\n", q.Len(), q.Cap())
	}
}
