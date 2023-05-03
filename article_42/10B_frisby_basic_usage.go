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

import "github.com/verdverm/frisby"

func main() {
	f := frisby.Create("Simplest test")
	f.Get("http://httpbin.org/get")
	f.Send()
	f.ExpectStatus(200)

	f = frisby.Create("Check HTTP code").Get("http://httpbin.org/status/321")
	f.Send()
	f.ExpectStatus(321)

	frisby.Global.PrintReport()
}
