// Seriál "Programovací jazyk Go"
//
// Třetí část
//
// Demonstrační příklad číslo 9:
//    Uživatelsky definovaná struktura.

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
