// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá devátá část
//    Trasování a profilování aplikací naprogramovaných v Go
//    https://www.root.cz/clanky/trasovani-a-profilovani-aplikaci-naprogramovanych-v-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z dvacáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_29/README.md

package slices

import "fmt"

func Slices() {
	var a [0]int
	s := a[:]

	fmt.Println(s)

	for i := 1; i <= 100000; i++ {
		s = append(s, i)
	}

	fmt.Println(s)
}
