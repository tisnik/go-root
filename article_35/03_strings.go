// Seriál "Programovací jazyk Go"
//
// Třicátá pátá část
//    Programovací jazyk Go pro skalní céčkaře (2.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/
//
// Demonstrační příklad číslo 3:
//     Práce s řetězci v jazyce Go.

package main

import "fmt"

func main() {
	var s1 string = "www.root.cz"
	var s2 string = ""
	var s3 string = "Hello\nworld!\n"
	var s4 string = "шщэюя"

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}
