// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 5:
//     Marshalling struktury/záznamu do JSONu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_13/05_json_marshal_struct.html

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

	user1JSON, _ := json.Marshal(user1)
	fmt.Println(string(user1JSON))

	user2JSON, _ := json.Marshal(user2)
	fmt.Println(string(user2JSON))
}
