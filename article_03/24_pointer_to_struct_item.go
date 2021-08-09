// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 24:
//    Ukazatel na položku struktury.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_03/24_pointer_to_struct_item.html

package main

import "fmt"

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	var u User

	u.id = 1
	u.name = "Pepek"
	u.surname = "Vyskoč"

	fmt.Println(u)

	var pName *string
	pName = &u.name

	fmt.Println(pName)
	fmt.Println(*pName)

	(*pName) = "Zdeněk"
	fmt.Println(*pName)
}
