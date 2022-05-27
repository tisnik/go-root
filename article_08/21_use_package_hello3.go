// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
//
// Seznam příkladů z osmé části:
//    https://github.com/tisnik/go-root/blob/master/article_08/README.md
//
// Demonstrační příklad číslo 21:
//    Import vlastního balíčku.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_08/21_use_package_hello3.html

package main

import (
	"hello3"
)

func main() {
	hello3.hello_world()
}
