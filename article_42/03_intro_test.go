// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá druhá část
//    Testování aplikací psaných v Go s využitím knihoven Goblin a Frisby
//    https://www.root.cz/clanky/testovani-aplikaci-psanych-v-go-s-vyuzitim-knihoven-goblin-a-frisby/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze čtyřicáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_42/README.md

package main

import (
	. "github.com/franela/goblin"
	"testing"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Adder", func() {
		g.It("Should add two numbers ", func() {
			g.Assert(1 + 1).Equal(2)
		})
		g.It("Should add two numbers", func() {
			g.Assert(1 + 1).Equal(5)
		})
		g.It("Should subtract two numbers")
		g.Xit("Should add two numbers, excluded ", func() {
			g.Assert(3 + 1).Equal(4)
		})
	})
}
