// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 16:
//    Modifikace návratové hodnoty funkce v sekci defer.

package main

import "fmt"

func funkce1() (i int) {
	i = 1
	return
}

func funkce2() (i int) {
	defer func() { i = 2 }()
	return 1
}

func funkce3() (i int) {
	defer func() { i += 2 }()
	return 1
}

func main() {
	fmt.Printf("Návratová hodnota funkce1: %d\n", funkce1())
	fmt.Printf("Návratová hodnota funkce2: %d\n", funkce2())
	fmt.Printf("Návratová hodnota funkce3: %d\n", funkce3())
}
