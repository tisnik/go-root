// Seriál "Programovací jazyk Go"
//
// Třetí část
//
// Demonstrační příklad číslo 19:
//    Mapa struktur.

package main

import "fmt"

// Key je uživatelsky definovaná datová struktura
type Key struct {
	id   uint32
	role string
}

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	m1 := make(map[Key]User)
	fmt.Println(m1)

	m1[Key{1, "admin"}] = User{
		id:      1,
		name:    "Pepek",
		surname: "Vyskoč"}

	m1[Key{2, "user"}] = User{
		id:      2,
		name:    "Josef",
		surname: "Vyskočil"}

	fmt.Println(m1)
}
