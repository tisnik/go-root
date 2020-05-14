// Seriál "Programovací jazyk Go"
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 5:
//     Marshalling struktury/záznamu do JSONu

package main

import (
	"encoding/json"
	"fmt"
)

// User1 je uživatelsky definovaná datová struktura s privátními atributy
type User1 struct {
	id      uint32
	name    string
	surname string
}

// User2 je uživatelsky definovaná datová struktura s viditelnými atributy
type User2 struct {
	ID      uint32
	Name    string
	Surname string
}

func main() {
	user1 := User1{
		1,
		"Pepek",
		"Vyskoč"}

	user2 := User2{
		1,
		"Pepek",
		"Vyskoč"}

	user1_json, _ := json.Marshal(user1)
	fmt.Println(string(user1_json))

	user2_json, _ := json.Marshal(user2)
	fmt.Println(string(user2_json))
}
