// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 11:
//    Uživatelsky definovaná struktura a její inicializace.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_03/11_better_struct_init.html

package main

import "fmt"

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	user1 := User{
		id:      1,
		name:    "Pepek",
		surname: "Vyskoč"}

	fmt.Println(user1)

	user1.id = 2
	user1.name = "Josef"
	user1.surname = "Vyskočil"

	fmt.Println(user1)
}
