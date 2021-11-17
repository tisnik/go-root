// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmnáctá část
//     Knihovny určené pro tvorbu testů v programovacím jazyce Go
//     https://www.root.cz/clanky/knihovny-urcene-pro-tvorbu-testu-v-programovacim-jazyce-go/
//
// Demonstrační příklad číslo 5:
//     Testovaný balíček.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_18/05_factorial_ogletest2/factorial.html

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
