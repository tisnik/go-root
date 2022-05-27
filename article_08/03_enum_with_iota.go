// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
//
// Seznam příkladů z osmé části:
//    https://github.com/tisnik/go-root/blob/master/article_08/README.md
//
// Demonstrační příklad číslo 3:
//    Náhrada za enum: běžné konstanty a identifikátor iota.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_08/03_enum_with_iota.html

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

func reservation(day Enum) {
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
