// Seriál "Programovací jazyk Go"
//
// Dvacátá pátá část
//     Ladění aplikací v Go s využitím GNU Debuggeru a debuggeru Delve
//     https://www.root.cz/clanky/ladeni-aplikaci-v-go-s-vyuzitim-gnu-debuggeru-a-debuggeru-delve/
//
// Demonstrační příklad číslo 3:
//     Práce se strukturami (záznamy)

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
		name:    "Linus",
		surname: "Torvalds"}

	fmt.Println(user1)

	user1.id = 2
	user1.name = "Steve"
	user1.surname = "Ballmer"

	fmt.Println(user1)
}
