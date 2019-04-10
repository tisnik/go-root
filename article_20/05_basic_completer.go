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
