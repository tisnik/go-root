// Seriál "Programovací jazyk Go"
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 30:
//    Dvojice vnořených smyček for a příkaz break.

package main

import "fmt"

func main() {
Exit:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%3d ", i*j)
			if i*j == 42 {
				fmt.Println("\nodpověď nalezena!\n")
				break Exit
			}
		}
		fmt.Println()
	}
}
