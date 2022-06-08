// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá část
//    Trasování a profilování aplikací naprogramovaných v Go (dokončení
//    https://www.root.cz/clanky/trasovani-a-profilovani-aplikaci-naprogramovanych-v-go-dokonceni/
//
// Seznam příkladů ze třicáté části:
//    https://github.com/tisnik/go-root/blob/master/article_30/README.md

package slices2

import "fmt"

func Slices() {
	var a [0]int
	s := a[:]

	for i := 1; i <= 1000000; i++ {
		s = append(s, i)
	}

	fmt.Printf("Length: %d\n", len(s))
}
