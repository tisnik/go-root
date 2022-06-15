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

import "github.com/verdverm/frisby"

func main() {
	f := frisby.Create("Simplest test").Post("http://httpbin.org/post").
		SetHeader("Content-Type", "application/json").SetHeader("Accept", "application/json").
		SetJson([]string{"item1", "item2", "item3"})
	f.Send()
	f.ExpectStatus(200)

	data := map[string]string{
		"text": "Hello **world**!",
	}
	f = frisby.Create("Markdown conversion").Post("https://api.github.com/markdown").SetJson(data)
	f.Send()
	f.ExpectStatus(200)

	frisby.Global.PrintReport()
}
