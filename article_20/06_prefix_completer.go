// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá část
//     Jazyk Go a aplikace s vlastním příkazovým řádkem
//     https://www.root.cz/clanky/jazyk-go-a-aplikace-s-vlastnim-prikazovym-radkem/
//
// Demonstrační příklad číslo 6:
//     Go-prompt: vstup s nápovědou, reakce na zadaný příkaz.
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
		println("Quitting")
		os.Exit(0)
	case "help":
		println("HELP:\nexit\nquit")
	default:
		println("Nothing happens")
	}
}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "help"},
		{Text: "exit"},
		{Text: "quit"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func main() {
	p := prompt.New(executor, completer)
	p.Run()
}
