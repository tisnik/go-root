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
// Demonstrační příklad číslo 3:
//     Testy pro balíček.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_18/03_factorial_oglematchers2/factorial_test.html

package factorial_test

import (
	"factorial"
	. "github.com/jacobsa/oglematchers"
	"testing"
)

func TestFactorialForZero(t *testing.T) {
	result := factorial.Factorial(0)
	m := Equals(1)
	if m.Matches(result) != nil {
		t.Errorf("Expected that 0! == 1, but got %d instead", result)
	}
}

func TestFactorialForOne(t *testing.T) {
	result := factorial.Factorial(1)
	m := Equals(1)
	if m.Matches(result) != nil {
		t.Errorf("Expected that 1! == 1, but got %d instead", result)
	}
}

func TestFactorialForSmallNumber(t *testing.T) {
	result := factorial.Factorial(5)
	m := AllOf(
		GreaterThan(10),
		LessThan(10000))
	if m.Matches(result) != nil {
		t.Errorf("Expected that 5! == is between 10..10000")
	}
}

func TestFactorialForSmallNumberNegative(t *testing.T) {
	result := factorial.Factorial(20)
	m := AllOf(
		GreaterThan(10),
		LessThan(10000))
	if m.Matches(result) != nil {
		t.Errorf("Expected that 20! == is between 10..10000")
	}
}

func TestFactorialForTen(t *testing.T) {
	result := factorial.Factorial(10)
	expected := int64(3628800)
	m := Equals(expected)
	if m.Matches(result) != nil {
		t.Errorf("Expected that 10! == %d, but got %d instead", expected, result)
	}
}

func TestFactorialForBigNumber(t *testing.T) {
	result := factorial.Factorial(20)
	m := GreaterThan(0)
	if m.Matches(result) != nil {
		t.Errorf("Expected that 20! > 0, but got negative number %d instead", result)
	}
}

func TestFactorialForEvenBiggerNumber(t *testing.T) {
	result := factorial.Factorial(30)
	m := GreaterThan(0)
	if m.Matches(result) != nil {
		t.Errorf("Expected that 30! > 0, but got negative number %d instead", result)
	}
}
