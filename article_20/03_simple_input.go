package main

import "github.com/c-bata/go-prompt"

func completer(in prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func main() {
	login := prompt.Input("Login: ", completer)
	password := prompt.Input("Password: ", completer)
	println(login)
	println(password)
}
