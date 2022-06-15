// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá druhá část
//    Testování aplikací psaných v Go s využitím knihoven Goblin a Frisby
//    https://www.root.cz/clanky/testovani-aplikaci-psanych-v-go-s-vyuzitim-knihoven-goblin-a-frisby/
//
// Seznam příkladů ze čtyřicáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_42/README.md

package main

import (
	"github.com/verdverm/frisby"
	"testing"
)

func TestHttpGet(t *testing.T) {
	f := frisby.Create("Simplest test").Get("http://httpbin.org/get")
	f.Send()
	f.ExpectStatus(200)
	err := f.Error()
	if err != nil {
		t.Error(err)
	}
}
