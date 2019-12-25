// Seriál "Programovací jazyk Go"
//
// Čtyřicátá osmá část
//
// Demonstrační příklad číslo 2:
//     Vytištění obsahu matice funkcí mat.Formatted

package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	zero := mat.NewDense(5, 6, nil)
	fmt.Println(mat.Formatted(zero))
}
