// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z osmé části:
//    https://github.com/tisnik/go-root/blob/master/article_08/README.md
//
// Demonstrační příklad číslo 20:
//    Import vlastního balíčku.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_08/20_use_package_hello2.html

package main

import (
	"hello2"
)

func main() {
	hello2.Hello_world()
}
