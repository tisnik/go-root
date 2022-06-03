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
// Demonstrační příklad číslo 1:
//     Testy pro balíček.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_18/01_factorial_testing/factorial_test.html

package factorial_test

import (
	"factorial"
	"testing"
)

func TestFactorialForZero(t *testing.T) {
	result := factorial.Factorial(0)
	if result != 1 {
		t.Errorf("Expected that 0! == 1, but got %d instead", result)
	}
}

func TestFactorialForOne(t *testing.T) {
	result := factorial.Factorial(1)
	if result != 1 {
		t.Errorf("Expected that 1! == 1, but got %d instead", result)
	}
}

func TestFactorialForSmallNumber(t *testing.T) {
	result := factorial.Factorial(5)
	if result <= 10 || result >= 10000 {
		t.Errorf("Expected that 5! == is between 10..10000")
	}
}

func TestFactorialForSmallNumberNegative(t *testing.T) {
	result := factorial.Factorial(20)
	if result <= 10 || result >= 10000 {
		t.Errorf("Expected that 20! == is between 10..10000")
	}
}

func TestFactorialForTen(t *testing.T) {
	result := factorial.Factorial(10)
	expected := int64(3628800)
	if result != expected {
		t.Errorf("Expected that 10! == %d, but got %d instead", expected, result)
	}
}

func TestFactorialForBigNumber(t *testing.T) {
	result := factorial.Factorial(20)
	if result <= 0 {
		t.Errorf("Expected that 20! > 0, but got negative number %d instead", result)
	}
}

func TestFactorialForEvenBiggerNumber(t *testing.T) {
	result := factorial.Factorial(30)
	if result <= 0 {
		t.Errorf("Expected that 30! > 0, but got negative number %d instead", result)
	}
}
