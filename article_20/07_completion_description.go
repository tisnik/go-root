// Seriál "Programovací jazyk Go"
//
// Dvacátá část
//
// Demonstrační příklad číslo 7:
//     Go-prompt: vylepšená nápověda.
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
	return
}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "help", Description: "show help with all commands"},
		{Text: "exit", Description: "quit the application"},
		{Text: "quit", Description: "quit the application"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func main() {
	p := prompt.New(executor, completer)
	p.Run()
}
