// Seriál "Programovací jazyk Go"
//
// Dvacátá pátá část
//     Ladění aplikací v Go s využitím GNU Debuggeru a debuggeru Delve
//     https://www.root.cz/clanky/ladeni-aplikaci-v-go-s-vyuzitim-gnu-debuggeru-a-debuggeru-delve/
//
// Demonstrační příklad číslo 4:
//     Rekurzivní výpočet faktoriálu

package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(4))
}
