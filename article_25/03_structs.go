// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá pátá část
//     Ladění aplikací v Go s využitím GNU Debuggeru a debuggeru Delve
//     https://www.root.cz/clanky/ladeni-aplikaci-v-go-s-vyuzitim-gnu-debuggeru-a-debuggeru-delve/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z dvacáté páté části:
//    https://github.com/tisnik/go-root/blob/master/article_25/README.md
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
