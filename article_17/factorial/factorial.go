// Seriál "Programovací jazyk Go"
//
// Balíček s funkcí pro výpočet faktoriálu

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
