// Seriál "Programovací jazyk Go"
//
// Dvacátá část část
//
// Demonstrační příklad číslo 4:
//     Go-prompt: vstup s nápovědou.

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
