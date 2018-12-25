// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 20:
//    Relační operátory.

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
