// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá pátá část
//     Testování Go aplikací s využitím knihovny GΩmega a frameworku Ginkgo
//     https://www.root.cz/clanky/testovani-go-aplikaci-s-vyuzitim-knihovny-gomega-mega-a-frameworku-ginkgo/
//
// Seznam příkladů ze čtyřicáté páté části:
//    https://github.com/tisnik/go-root/blob/master/article_45/README.md

package main

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestCollecions(t *testing.T) {
	g := NewGomegaWithT(t)

	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	g.Expect(m).To(HaveLen(3))
	g.Expect(m).To(ContainElement(1))
	g.Expect(m).To(Not(ContainElement(42)))
	g.Expect(m).To(HaveKey("two"))
	g.Expect(m).To(Not(HaveKey("forty two")))
}
