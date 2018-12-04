// Seriál "Programovací jazyk Go"
//
// Třetí část
//
// Demonstrační příklad číslo 29:
//    Datový typ "funkce s návratovou hodnotou"

package main

import "fmt"

type no_param_function func() int
type two_int_param_function func(int, int) int

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
	var a no_param_function
	var b two_int_param_function

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
