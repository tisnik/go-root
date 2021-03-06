// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá pátá část
//     Testování Go aplikací s využitím knihovny GΩmega a frameworku Ginkgo
//     https://www.root.cz/clanky/testovani-go-aplikaci-s-vyuzitim-knihovny-gomega-mega-a-frameworku-ginkgo/

package main

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestFactorialForZero(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(Factorial(0)).To(BeNumerically("==", 1))
}

func TestFactorialForOne(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(Factorial(1)).To(BeNumerically("==", 1))
}

func TestFactorialForTen(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(Factorial(10)).To(BeNumerically("==", 3628800))
}

func TestFactorialForSmallNumber(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(Factorial(5)).To(SatisfyAll(BeNumerically(">=", 10), BeNumerically("<=", 10000)))
}
