// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 9:
//    Uživatelsky definovaná struktura.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_03/09_struct.html

package main

import "fmt"

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	var user1 User

	fmt.Println(user1)

	user1.id = 1
	user1.name = "Pepek"
	user1.surname = "Vyskoč"

	fmt.Println(user1)
}
