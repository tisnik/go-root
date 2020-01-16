// Seriál "Programovací jazyk Go"
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 23:
//    Ukazatel na strukturu.

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

	var pUser *User
	pUser = &u

	fmt.Println(pUser)
	fmt.Println(*pUser)

	(*pUser).id = 10000
	fmt.Println(*pUser)

	pUser.id = 20000
	fmt.Println(*pUser)
}
