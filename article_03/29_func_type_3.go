// Seriál "Programovací jazyk Go"
//
// Třetí část
//
// Demonstrační příklad číslo 29:
//    Datový typ "funkce s návratovou hodnotou"

package main

import "fmt"

type noParamFunction func() int
type twoIntParamFunction func(int, int) int

func funkce1() int {
	return 42
}

func funkce2() int {
	return -1
}

func funkce3(x int, y int) int {
	return x + y
}

func funkce4(x int, y int) int {
	return x * y
}

func main() {
	var a noParamFunction
	var b twoIntParamFunction

	fmt.Println(a)
	fmt.Println(b)

	a = funkce1
	fmt.Println(a)
	fmt.Println(a())

	a = funkce2
	fmt.Println(a)
	fmt.Println(a())

	b = funkce3
	fmt.Println(b)
	fmt.Println(b(10, 20))

	b = funkce4
	fmt.Println(b)
	fmt.Println(b(10, 20))
}
