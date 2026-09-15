// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá šestá část
//    Kooperace mezi kódem psaným v Go a C: cgo
//    https://www.root.cz/clanky/kooperace-mezi-kodem-psanym-v-go-a-c-cgo/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Demonstrační příklad číslo 8:
//    Volání nativní funkce s předáním hodnoty s plovoucí řádovou čárkou.
//
// Seznam demonstračních příkladů ze třicáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_36/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_36/08_math.html
//

package main

// #include <math.h>
import "C"
import "fmt"

func main() {
	x := 3.1415927 / 6.0
	y := float32(C.sinf(C.float(x)))
	fmt.Printf("%v\n", y)
}
