// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmnáctá část
//     Testování aplikací naprogramovaných v jazyce Go
//     https://www.root.cz/clanky/testovani-aplikaci-naprogramovanych-v-jazyce-go/
//
// Jednotkové testy pro výpočet faktoriálu.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_17/factorial/factorial_test.html

package factorial_test

import (
	"factorial"
	"testing"
)

func TestFactorialForZero(t *testing.T) {
	result := factorial.Factorial(0)
	if result != 1 {
		t.Errorf("Expected 0! == 1, but got %d instead", result)
	}
}

func TestFactorialForTen(t *testing.T) {
	result := factorial.Factorial(10)
	expected := int64(3628800)
	if result != expected {
		t.Errorf("Expected 0! == %d, but got %d instead", expected, result)
	}
}
