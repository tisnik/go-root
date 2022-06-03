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
// Demonstrační příklad číslo 4:
//     Testy pro balíček.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_18/04_factorial_ogletest/factorial_test.html

package factorial_test

import (
	"factorial"
	"github.com/jacobsa/oglematchers"
	"github.com/jacobsa/ogletest"
	"testing"
)

func TestOgletest(t *testing.T) {
	ogletest.RunTests(t)
}

type FactorialTest struct{}

func init() {
	ogletest.RegisterTestSuite(&FactorialTest{})
}

func (t *FactorialTest) FactorialForZero() {
	result := factorial.Factorial(0)
	ogletest.ExpectThat(result, oglematchers.Equals(1))
}

func (t *FactorialTest) FactorialForOne() {
	result := factorial.Factorial(1)
	ogletest.ExpectThat(result, oglematchers.Equals(1))
}

func (t *FactorialTest) TestFactorialSmallNumber() {
	result := factorial.Factorial(5)
	ogletest.ExpectThat(result, oglematchers.AllOf(
		oglematchers.GreaterThan(10),
		oglematchers.LessThan(10000)))
}

func (t *FactorialTest) TestFactorialSmallNumberNegative() {
	result := factorial.Factorial(20)
	ogletest.ExpectThat(result, oglematchers.AllOf(
		oglematchers.GreaterThan(10),
		oglematchers.LessThan(10000)))
}

func (t *FactorialTest) TestFactorialForTen() {
	result := factorial.Factorial(10)
	expected := int64(3628800)
	ogletest.ExpectThat(result, oglematchers.Equals(expected))
}

func (t *FactorialTest) TestFactorialBigNumber() {
	result := factorial.Factorial(20)
	ogletest.ExpectThat(result, oglematchers.GreaterThan(0))
}

func (t *FactorialTest) TestFactorialEvenBiggerNumber() {
	result := factorial.Factorial(30)
	ogletest.ExpectThat(result, oglematchers.GreaterThan(0))
}
