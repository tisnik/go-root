// Seriál "Programovací jazyk Go"
//
// Třicátá pátá část
//    Programovací jazyk Go pro skalní céčkaře (2.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/
//
// Demonstrační příklad číslo 5:
//    Alternativa k měnitelným řetězcům.

package main

import "fmt"

func main() {
	var s []byte = []byte("www.root.cz")

	fmt.Println(string(s))
	s[3] = '*'
	s[8] = '*'
	fmt.Println(string(s))
}
