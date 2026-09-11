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
//
// Demonstrační příklad číslo 9A:
//    Získání části řetězce.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_35/09_substring_B.html
//

package main

import "fmt"

func main() {
	s := "шщэюя"

	fmt.Println(s[2:4])
	fmt.Println(s[2:5])
	fmt.Println(s[2:6])

	fmt.Println()
	fmt.Println(s[3:4])
	fmt.Println(s[3:5])
	fmt.Println(s[3:6])
}
