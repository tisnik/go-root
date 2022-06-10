// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá pátá část
//    Programovací jazyk Go pro skalní céčkaře (2.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/
//
// Seznam příkladů ze třicáté páté části:
//    https://github.com/tisnik/go-root/blob/master/article_35/README.md

package main

import "fmt"

func main() {
	s1 := "Hello world!"

	s2 := s1[0:4]
	fmt.Println(s2)
}
