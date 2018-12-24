// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 10:
//    Zjištění pořadí volání funkcí deklarovaných v defer().

package main

import "fmt"

func on_finish(i int) {
	fmt.Printf("Defer #%2d\n", i)
}

func main() {
	for i := 0; i <= 10; i++ {
		defer on_finish(i)
	}
	fmt.Println("Finishing main() function")
}
