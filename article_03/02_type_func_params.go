// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 2:
//    Kontrola parametrů funkcí.

package main

import "fmt"

// ID of user
type ID uint32

// Name of user
type Name string

// Surname of user
type Surname string

func registerUser(id ID, name Name, surname Surname) {
	fmt.Printf("Registering: %d %s %s", id, name, surname)
}

func main() {
	var i ID = 1
	var n Name = "Jan"
	var s Surname = "Novák"

	registerUser(i, n, s)
}
