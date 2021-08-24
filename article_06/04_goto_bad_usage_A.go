// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 4:
//    Špatné použití příkazu goto.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_06/04_goto_bad_usage_A.html

package main

import "fmt"

func main() {
	goto Next
	i := 10
Next:
	fmt.Printf("%2d\n", i)
}
