// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmnáctá část
//     Knihovny určené pro tvorbu testů v programovacím jazyce Go
//     https://www.root.cz/clanky/knihovny-urcene-pro-tvorbu-testu-v-programovacim-jazyce-go/
//
// Seznam příkladů z osmnácté části:
//    https://github.com/tisnik/go-root/blob/master/article_18/README.md
//
// Demonstrační příklad číslo 2:
//     Testovaný balíček.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_18/02_factorial_oglematchers/factorial.html

package factorial

// Factorial computes factorial for given input using recurrence relation formula
func Factorial(n int64) int64 {
	switch {
	case n < 0:
		return 1
	case n == 0:
		return 1
	default:
		return n * Factorial(n-1)
	}
}
