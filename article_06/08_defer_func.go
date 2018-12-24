// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 8:
//    Použití příkazu defer.

package main

import "fmt"

func main() {
	defer (func() { fmt.Println("Finished") })()

	for i := 10; i >= 0; i-- {
		fmt.Printf("%2d\n", i)
	}
	fmt.Println("Finishing main() function")
}
