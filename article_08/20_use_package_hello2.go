// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
//
// Demonstrační příklad číslo 20:
//    Import vlastního balíčku.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_07/20_use_package_hello2.html

package main

import (
	"hello2"
)

func main() {
	hello2.Hello_world()
}
