// Seriál "Programovací jazyk Go"
//
// Osmnáctá část
//     Knihovny určené pro tvorbu testů v programovacím jazyce Go
//     https://www.root.cz/clanky/knihovny-urcene-pro-tvorbu-testu-v-programovacim-jazyce-go/
//
// Demonstrační příklad číslo 9:
//     Testy pro balíček.

package factorial

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFactorialForZero(t *testing.T) {
	convey.Convey("0! should be equal 1", t, func() {
		convey.So(Factorial(0), convey.ShouldEqual, 1)
	})
}

func TestFactorialForOne(t *testing.T) {
	convey.Convey("1! should be equal 1", t, func() {
		convey.So(Factorial(1), convey.ShouldEqual, 1)
	})
}

func TestFactorialForSmallNumber(t *testing.T) {
	convey.Convey("5! should be between 1 and 10000", t, func() {
		convey.So(Factorial(5), convey.ShouldBeBetween, 1, 10000)
	})
}

func TestFactorialForSmallNumberNegative(t *testing.T) {
	convey.Convey("20! should be between 1 and 10000", t, func() {
		convey.So(Factorial(20), convey.ShouldBeBetween, 1, 10000)
	})
}

func TestFactorialForTen(t *testing.T) {
	convey.Convey("10! should be equal to 3628800", t, func() {
		convey.So(Factorial(10), convey.ShouldEqual, 3628800)
	})
}

func TestFactorialForBigNumber(t *testing.T) {
	convey.Convey("20! should be greater than zero", t, func() {
		convey.So(Factorial(20), convey.ShouldBeGreaterThan, 0)
	})
}

func TestFactorialForEvenBiggerNumber(t *testing.T) {
	convey.Convey("30! should be greater than zero", t, func() {
		convey.So(Factorial(30), convey.ShouldBeGreaterThan, 0)
	})
}
