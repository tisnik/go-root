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
	"github.com/verdverm/frisby"
	"testing"
)

func TestHttpGet200Ok(t *testing.T) {
	f := frisby.Create("Simplest test for HTTP 200 OK").Get("http://httpbin.org/get")
	f.Send()
	f.ExpectStatus(200)
	err := f.Error()
	if err != nil {
		t.Error(err)
	}
}

func TestHttpGet400Ok(t *testing.T) {
	f := frisby.Create("Simplest test for HTTP 400 Bad Request").Get("http://httpbin.org/get")
	f.Send()
	f.ExpectStatus(400)
	err := f.Error()
	if err != nil {
		t.Error(err)
	}
}
