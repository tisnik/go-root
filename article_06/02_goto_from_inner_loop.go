// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 2:
//    Dvojice vnořených smyček for a příkaz goto pro jejich opuštění.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%3d ", i*j)
			if i*j == 42 {
				fmt.Println("\nodpověď nalezena!\n")
				goto Exit
			}
		}
		fmt.Println()
	}
Exit:
}
