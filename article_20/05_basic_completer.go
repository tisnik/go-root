// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá část
//     Jazyk Go a aplikace s vlastním příkazovým řádkem
//     https://www.root.cz/clanky/jazyk-go-a-aplikace-s-vlastnim-prikazovym-radkem/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z dvacáté části:
//    https://github.com/tisnik/go-root/blob/master/article_20/README.md
//
// Demonstrační příklad číslo 5:
//     Go-prompt: vstup s nápovědou.
//
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_20/05_basic_completer.html
//

package main

import (
	"github.com/c-bata/go-prompt"
	"os"
)

func executor(t string) {
	switch t {
	case "exit":
		fallthrough
	case "quit":
		os.Exit(0)
	case "help":
		println("HELP:\nexit\nquit")
	default:
		println("Nothing happens")
	}
}

func completer(in prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "help"},
		{Text: "exit"},
		{Text: "quit"},
	}
}

func main() {
	p := prompt.New(executor, completer)
	p.Run()
}
