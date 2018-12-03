// Seriál "Programovací jazyk Go"
//
// Třetí část
//
// Demonstrační příklad číslo 12:
//    Porovnání struktur.

package main

import "fmt"

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

	var user2 User
	user2.id = 1
	user2.name = "Pepek"
	user2.surname = "Vyskoč"

	user3 := User{
		id:      1,
		name:    "Josef",
		surname: "Vyskočil"}

	fmt.Println(user1 == user2)
	fmt.Println(user1 == user3)
}
