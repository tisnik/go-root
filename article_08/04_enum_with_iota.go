// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
//
// Demonstrační příklad číslo 4:
//    Převod "výčtu" na řetězec.

package main

import "fmt"

// Enum je užitevatelsky definovaný datový typ mj. i pro dny v týdnu
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

func reservation(day Enum) {
	fmt.Printf("Reservation for %s\n", day.String())
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
