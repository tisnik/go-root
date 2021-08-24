// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 20:
//    Relační operátory.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_06/20_relational_operators.html

package main

import "fmt"

func main() {
	x := 42
	y := 0

	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x == y)
	fmt.Println(x >= y)
	fmt.Println(x > y)
	fmt.Println(x != y)

	fmt.Println()

	fx := 1e10
	fy := -2.3e56

	fmt.Println(fx < fy)
	fmt.Println(fx <= fy)
	fmt.Println(fx == fy)
	fmt.Println(fx >= fy)
	fmt.Println(fx > fy)
	fmt.Println(fx != fy)

	fmt.Println()

	bx := true
	by := false

	fmt.Println(bx == by)
	fmt.Println(bx != by)

	fmt.Println()

	cx := 1 + 1i
	cy := 0 + 0i

	fmt.Println(cx == cy)
	fmt.Println(cx != cy)

	fmt.Println()

	sx := "Hello"
	sy := "World"

	fmt.Println(sx < sy)
	fmt.Println(sx <= sy)
	fmt.Println(sx == sy)
	fmt.Println(sx >= sy)
	fmt.Println(sx > sy)
	fmt.Println(sx != sy)

	fmt.Println()

}
