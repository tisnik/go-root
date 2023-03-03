// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmnáctá část
//     Knihovny určené pro tvorbu testů v programovacím jazyce Go
//     https://www.root.cz/clanky/knihovny-urcene-pro-tvorbu-testu-v-programovacim-jazyce-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z osmnácté části:
//    https://github.com/tisnik/go-root/blob/master/article_18/README.md
//
// Demonstrační příklad číslo 2:
//     Testy pro balíček.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_18/02_factorial_oglematchers/factorial_test.html

package factorial_test

import (
	"factorial"
	"github.com/jacobsa/oglematchers"
	"testing"
)

func TestFactorialForZero(t *testing.T) {
	result := factorial.Factorial(0)
	m := oglematchers.Equals(1)
	if m.Matches(result) != nil {
		t.Errorf("Expected that 0! == 1, but got %d instead", result)
	}
}

func TestFactorialForOne(t *testing.T) {
	result := factorial.Factorial(1)
	m := oglematchers.Equals(1)
	if m.Matches(result) != nil {
		t.Errorf("Expected that 1! == 1, but got %d instead", result)
	}
}

func TestFactorialForSmallNumber(t *testing.T) {
	result := factorial.Factorial(5)
	m := oglematchers.AllOf(
		oglematchers.GreaterThan(10),
		oglematchers.LessThan(10000))
	if m.Matches(result) != nil {
		t.Errorf("Expected that 5! == is between 10..10000")
	}
}

func TestFactorialForSmallNumberNegative(t *testing.T) {
	result := factorial.Factorial(20)
	m := oglematchers.AllOf(
		oglematchers.GreaterThan(10),
		oglematchers.LessThan(10000))
	if m.Matches(result) != nil {
		t.Errorf("Expected that 20! == is between 10..10000")
	}
}

func TestFactorialForTen(t *testing.T) {
	result := factorial.Factorial(10)
	expected := int64(3628800)
	m := oglematchers.Equals(expected)
	if m.Matches(result) != nil {
		t.Errorf("Expected that 10! == %d, but got %d instead", expected, result)
	}
}

func TestFactorialForBigNumber(t *testing.T) {
	result := factorial.Factorial(20)
	m := oglematchers.GreaterThan(0)
	if m.Matches(result) != nil {
		t.Errorf("Expected that 20! > 0, but got negative number %d instead", result)
	}
}

func TestFactorialForEvenBiggerNumber(t *testing.T) {
	result := factorial.Factorial(30)
	m := oglematchers.GreaterThan(0)
	if m.Matches(result) != nil {
		t.Errorf("Expected that 30! > 0, but got negative number %d instead", result)
	}
}
