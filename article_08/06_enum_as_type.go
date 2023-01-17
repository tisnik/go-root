// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z osmé části:
//    https://github.com/tisnik/go-root/blob/master/article_08/README.md
//
// Demonstrační příklad číslo 6:
//    Výčet jako datový typ: nemožnost použití operátoru ++.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_08/06_enum_as_type.html

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

// Den je uživatelsky definovaná datová struktura reprezentující jeden den v týdnu
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
