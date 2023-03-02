// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmnáctá část
//     Testování aplikací naprogramovaných v jazyce Go
//     https://www.root.cz/clanky/testovani-aplikaci-naprogramovanych-v-jazyce-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze sedmnácté části:
//    https://github.com/tisnik/go-root/blob/master/article_17/README.md
//
// Balíček s funkcí pro výpočet faktoriálu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_17/factorial/factorial.html

package factorial

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
