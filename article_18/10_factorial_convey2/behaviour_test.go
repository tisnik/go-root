// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmnáctá část
//     Knihovny určené pro tvorbu testů v programovacím jazyce Go
//     https://www.root.cz/clanky/knihovny-urcene-pro-tvorbu-testu-v-programovacim-jazyce-go/
//
// Demonstrační příklad číslo 10:
//     Testy pro balíček.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_18/10_factorial_convey2/behaviour_test.html

package factorial

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFactorialForZero(t *testing.T) {
	Convey("0! should be equal 1", t, func() {
		So(Factorial(0), ShouldEqual, 1)
	})
}

func TestFactorialForOne(t *testing.T) {
	Convey("1! should be equal 1", t, func() {
		So(Factorial(1), ShouldEqual, 1)
	})
}

func TestFactorialForSmallNumber(t *testing.T) {
	Convey("5! should be between 1 and 10000", t, func() {
		So(Factorial(5), ShouldBeBetween, 1, 10000)
	})
}

func TestFactorialForSmallNumberNegative(t *testing.T) {
	Convey("20! should be between 1 and 10000", t, func() {
		So(Factorial(20), ShouldBeBetween, 1, 10000)
	})
}

func TestFactorialForTen(t *testing.T) {
	Convey("10! should be equal to 3628800", t, func() {
		So(Factorial(10), ShouldEqual, 3628800)
	})
}

func TestFactorialForBigNumber(t *testing.T) {
	Convey("20! should be greater than zero", t, func() {
		So(Factorial(20), ShouldBeGreaterThan, 0)
	})
}

func TestFactorialForEvenBiggerNumber(t *testing.T) {
	Convey("30! should be greater than zero", t, func() {
		So(Factorial(30), ShouldBeGreaterThan, 0)
	})
}
