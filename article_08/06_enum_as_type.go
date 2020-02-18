// Seriál "Programovací jazyk Go"
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
//
// Demonstrační příklad číslo 6:
//    Výčet jako datový typ: nemožnost použití operátoru ++.

package main

import "fmt"

type Enum int

const (
	Pondeli Enum = iota
	Utery
	Streda
	Ctvrtek
	Patek
	Sobota
	Nedele
)

type Den struct {
	X Enum
}

func (day Enum) String() string {
	days := []string{
		"Pondeli",
		"Utery",
		"Streda",
		"Ctvrtek",
		"Patek",
		"Sobota",
		"Nedele"}
	if day < Pondeli || day > Nedele {
		return "Unknown day"
	}
	return days[day]
}

func reservation(day Den) {
	fmt.Printf("Reservation for %s\n", day)
}

func main() {
	reservation(Den{Pondeli})
	reservation(Den{Sobota})
	reservation(Den{Nedele})

	day := Den{Pondeli}
	day++
	reservation(day)
}
