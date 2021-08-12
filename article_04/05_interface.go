// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 5:
//    Definice trojice jednoduchých rozhraní.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_04/05_interface.html

package main

// Shape je uživatelsky definovaná datová struktura
// představující geometrický tvar
type Shape interface {
}

// OpenShape je uživatelsky definovaná datová struktura
// představující otevřené tvary (úsečka, oblouk, křivka)
type OpenShape interface {
	length() float64
}

// ClosedShape je uživatelsky definovaná datová struktura
// představující uzavřené tvary (obdélník, kružnice, elipsa)
type ClosedShape interface {
	area() float64
	perimeter() float64
}

func main() {
}
