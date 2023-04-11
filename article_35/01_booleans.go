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
// Demonstrační příklad číslo 1:
//     Pravdivostní hodnoty v jazyce Go.

package main

func main() {
	if nil {
		println("true")
	}

	if !nil {
		println("false")
	}

	if 1 {
		println("true")
	}

	if 0 {
		println("false")
	}

	if 1.5 {
		println("true")
	}

	if 0.0 {
		println("false")
	}

	var b1 bool = nil

	if b1 {
		println("true")
	}

	if !b1 {
		println("false")
	}
}
