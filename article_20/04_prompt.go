// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá část
//     Jazyk Go a aplikace s vlastním příkazovým řádkem
//     https://www.root.cz/clanky/jazyk-go-a-aplikace-s-vlastnim-prikazovym-radkem/
//
// Demonstrační příklad číslo 4:
//     Go-prompt: vstup s nápovědou.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_20/04_prompt.html
//

package main

import "github.com/c-bata/go-prompt"

func executor(t string) {
	println(t)
}

func completer(in prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func main() {
	p := prompt.New(executor, completer)
	p.Run()
}
