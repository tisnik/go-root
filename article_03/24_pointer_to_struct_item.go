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

	var pName *string
	pName = &u.name

	fmt.Println(pName)
	fmt.Println(*pName)

	(*pName) = "Zdeněk"
	fmt.Println(*pName)
}
