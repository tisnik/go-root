// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 9:
//    Použití příkazu defer, zavolání funkce s parametrem.

package main

import "fmt"

func on_finish(message string) {
	fmt.Println(message)
}

func main() {
	defer on_finish("Finished")

	for i := 10; i >= 0; i-- {
		fmt.Printf("%2d\n", i)
	}
	fmt.Println("Finishing main() function")
}
