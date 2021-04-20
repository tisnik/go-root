// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá pátá část
//    Programovací jazyk Go pro skalní céčkaře (2.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/
//
// Demonstrační příklad číslo 4:
//     Řetězce jsou v Go neměnitelné.

package main

import "fmt"

func main() {
	var s string = "www.root.cz"

	fmt.Println(s)
	s[3] = '*'
	s[8] = '*'
	fmt.Println(s)
}
