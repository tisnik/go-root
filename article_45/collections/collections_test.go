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

func TestCollecions(t *testing.T) {
	g := NewGomegaWithT(t)
	c := []int{1, 2, 3, 4, 5}
	g.Expect(c).To(HaveLen(5))
	g.Expect(c).To(HaveCap(5))
	g.Expect(c).To(ContainElement(1))
	g.Expect(c).To(Not(ContainElement(42)))
}
