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
	frisby.Create("Cookies check").
		Get("http://httpbin.org/cookies").
		SetCookie("foo", "bar").
		Send().
		ExpectStatus(200).
		ExpectJson("cookies.foo", "bar").
		PrintBody()

	frisby.Global.PrintReport()
}
