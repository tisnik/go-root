// Seriál "Programovací jazyk Go"
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
//
// Demonstrační příklad číslo 2:
//    Náhrada za enum: běžné konstanty a identifikátor iota.

package main

import "fmt"

type Enum int

// Konstanty reprezentující dny v týdnu
const (
	Pondeli Enum = iota
	Utery
	Streda
	Ctvrtek
	Patek
	Sobota
	Nedele
)

func reservation(day int) {
	fmt.Printf("Reservation for %d\n", day)
}

func main() {
	reservation(Pondeli)
	reservation(Sobota)
	reservation(Nedele)

	reservation(3)

	day := Pondeli
	day--
	reservation(day)
}
