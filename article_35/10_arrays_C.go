// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá pátá část
//    Programovací jazyk Go pro skalní céčkaře (2.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze třicáté páté části:
//    https://github.com/tisnik/go-root/blob/master/article_35/README.md

package main

import "fmt"

func main() {
	var a [3][4]int

	fmt.Printf("Pole pred upravou: %v\n", a)

	for i := 0; i < 12; i++ {
		a[0][i] = i
	}

	fmt.Printf("Pole po uprave:    %v\n", a)
}
