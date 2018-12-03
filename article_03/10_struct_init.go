// Seriál "Programovací jazyk Go"
//
// Třetí část
//
// Demonstrační příklad číslo 10:
//    Uživatelsky definovaná struktura a její inicializace.

package main

import "fmt"

type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	user1 := User{
		1,
		"Pepek",
		"Vyskoč"}

	fmt.Println(user1)

	user1.id = 2
	user1.name = "Josef"
	user1.surname = "Vyskočil"

	fmt.Println(user1)
}
