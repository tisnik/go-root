// Seriál "Programovací jazyk Go"
//
// Čtyřicátá druhá část
//    Testování aplikací psaných v Go s využitím knihoven Goblin a Frisby
//    https://www.root.cz/clanky/testovani-aplikaci-psanych-v-go-s-vyuzitim-knihoven-goblin-a-frisby/

package main

import "github.com/verdverm/frisby"

func main() {
	f := frisby.Create("Simplest test")
	f.Get("http://httpbin.org/get")
	f.Send()
	f.ExpectStatus(200)
	f.PrintReport()

	f = frisby.Create("Check HTTP code").Get("http://httpbin.org/status/321")
	f.Send()
	f.ExpectStatus(321)
	f.PrintReport()

	frisby.Global.PrintReport()
}
