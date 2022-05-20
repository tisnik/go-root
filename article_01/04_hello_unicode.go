// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Seznam příkladů z první části:
//    https://github.com/tisnik/go-root/blob/master/article_01/README.md
//
// Demonstrační příklad číslo 4:
//    Použití Unicode ve zdrojovém kódu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_01/04_hello_unicode.html

package main

import "fmt"

func main() {
	fmt.Println("╭─────────────────────╮")
	fmt.Println("│ příλiš žλuťΩučký kůň│")
	fmt.Println("╰─────────────────────╯")
}
