// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z devadesáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_92/README.md

package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/arraylist"
	"github.com/daichi-m/go18ds/utils"
)

func main() {
	l := arraylist.New[string]("zoo", "aleph", "foo", "bar", "baz")
	fmt.Println(l)
	fmt.Println()

	l.Sort(utils.StringComparator)
	fmt.Println(l)
}
