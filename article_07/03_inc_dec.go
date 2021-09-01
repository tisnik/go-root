// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Demonstrační příklad číslo 3:
//    Operátory ++ a --.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_07/03_inc_dec.html

package main

import "fmt"

func main() {
	x := 1
	fmt.Printf("x = %d\n", x)

	x++
	fmt.Printf("x = %d\n", x)

	x--
	fmt.Printf("x = %d\n", x)

	px := &x
	fmt.Printf("x = %d\n", *px)

	*px++
	fmt.Printf("x = %d\n", *px)

	*px--
	fmt.Printf("x = %d\n", *px)

	y := 3.14
	fmt.Printf("y = %f\n", y)
	y++
	fmt.Printf("y = %f\n", y)

	z := 1 + 2i
	fmt.Printf("z = %f\n", z)
	z++
	fmt.Printf("z = %f\n", z)
}
