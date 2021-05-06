// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá druhá část
//    Testování aplikací psaných v Go s využitím knihoven Goblin a Frisby
//    https://www.root.cz/clanky/testovani-aplikaci-psanych-v-go-s-vyuzitim-knihoven-goblin-a-frisby/

package main

import "github.com/verdverm/frisby"

func main() {
	frisby.Create("Cookies check").
		Get("http://httpbin.org/cookies").
		SetCookie("foo", "bar").
		Send().
		ExpectStatus(200).
		ExpectJson("cookies.foo", "bar").
		PrintBody()

	frisby.Global.PrintReport()
}
