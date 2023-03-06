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
// Demonstrační příklad číslo 3:
//     Go-prompt: vstup dat.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_20/03_simple_input.html
//

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
