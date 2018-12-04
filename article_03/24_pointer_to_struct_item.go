// Seriál "Programovací jazyk Go"
//
// Třetí část
//
// Demonstrační příklad číslo 24:
//    Ukazatel na položku struktury.

package main

import "fmt"

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

	var p_n *string
	p_n = &u.name

	fmt.Println(p_n)
	fmt.Println(*p_n)

	(*p_n) = "Zdeněk"
	fmt.Println(*p_n)
}
