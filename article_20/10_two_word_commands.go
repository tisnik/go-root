// Seriál "Programovací jazyk Go"
//
// Dvacátá část část
//
// Demonstrační příklad číslo 10:
//     Víceslovní příkazy.

package main

import (
	"github.com/c-bata/go-prompt"
	"os"
	"strings"
)

func executor(t string) {
	switch t {
	case "exit":
		fallthrough
	case "quit":
		os.Exit(0)
	case "help":
		println("HELP:\nexit\nquit")
	case "user add":
		println("Adding user")
	case "user del":
		println("Deleting user")
	default:
		println("Nothing happens")
	}
	return
}

func completer(in prompt.Document) []prompt.Suggest {
	blocks := strings.Split(in.TextBeforeCursor(), " ")

	s := []prompt.Suggest{
		{Text: "help", Description: "show help with all commands"},
		{Text: "user", Description: "add/delete user"},
		{Text: "ls", Description: "list users/storages/computers"},
		{Text: "list", Description: "list users/storages/computers"},
		{Text: "exit", Description: "quit the application"},
		{Text: "quit", Description: "quit the application"},
	}

	user_s := []prompt.Suggest{
		{Text: "add", Description: "add new user"},
		{Text: "assign", Description: "assign a role to user"},
		{Text: "del", Description: "delete user"},
	}

	list_s := []prompt.Suggest{
		{Text: "users", Description: "show list of all users"},
		{Text: "logs", Description: "show list of all logs"},
		{Text: "storages", Description: "show list of all storages"},
		{Text: "computers", Description: "show list of all computers"},
	}

	empty_s := []prompt.Suggest{}

	if len(blocks) == 2 {
		switch blocks[0] {
		case "user":
			return prompt.FilterHasPrefix(user_s, blocks[1], true)
		case "ls":
			fallthrough
		case "list":
			return prompt.FilterHasPrefix(list_s, blocks[1], true)
		default:
			return empty_s
		}
	}
	return prompt.FilterHasPrefix(s, blocks[0], true)
}

func main() {
	p := prompt.New(executor, completer)
	p.Run()
}
