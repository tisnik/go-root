// Seriál "Programovací jazyk Go"
//
// Čtyřicátá druhá část
//    Testování aplikací psaných v Go s využitím knihoven Goblin a Frisby
//    https://www.root.cz/clanky/testovani-aplikaci-psanych-v-go-s-vyuzitim-knihoven-goblin-a-frisby/

package main

import "github.com/verdverm/frisby"

func main() {
	f := frisby.Create("Headers check").Get("http://httpbin.org/get")
	f.Send()
	f.ExpectStatus(200)
	f.ExpectHeader("Server", "apache")

	frisby.Global.PrintReport()
}
