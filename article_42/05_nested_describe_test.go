// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá druhá část
//    Testování aplikací psaných v Go s využitím knihoven Goblin a Frisby
//    https://www.root.cz/clanky/testovani-aplikaci-psanych-v-go-s-vyuzitim-knihoven-goblin-a-frisby/

package main

import (
	. "github.com/franela/goblin"
	"testing"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Adder", func() {
		g.Describe("Positive numbers", func() {
			g.It("Should add two numbers ", func() {
				g.Assert(1 + 1).Equal(2)
			})
			g.It("Should add two numbers", func() {
				g.Assert(2 + 2).Equal(4)
			})
			g.It("Should add two numbers", func() {
				g.Assert(10 + 20).Equal(30)
			})
		})
		g.Describe("Negative numbers", func() {
			g.It("Should add two numbers ", func() {
				g.Assert(1 + -1).Equal(0)
			})
			g.It("Should add two numbers", func() {
				g.Assert(2 + -4).Equal(-2)
			})
			g.It("Should add two numbers", func() {
				g.Assert(10 + -20).Equal(-10)
			})
		})
	})
}
