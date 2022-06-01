// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Seznam příkladů z třinácté části:
//    https://github.com/tisnik/go-root/blob/master/article_13/README.md
//
// Demonstrační příklad číslo 8:
//     Marshalling map struktur/záznamů do JSONu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_13/08_json_marshal_map_of_structs.html

package main

import (
	"encoding/json"
	"fmt"
)

// Key reprezentuje klíč použitý pro identifikaci uživatele
type Key struct {
	ID   uint32
	Role string
}

// User je uživatelsky definovaná datová struktura
type User struct {
	ID      uint32
	Name    string
	Surname string
}

func main() {
	m1 := make(map[Key]User)

	m1[Key{1, "admin"}] = User{
		ID:      1,
		Name:    "Pepek",
		Surname: "Vyskoč"}

	m1[Key{2, "user"}] = User{
		ID:      2,
		Name:    "Josef",
		Surname: "Vyskočil"}

	m1JSON, _ := json.Marshal(m1)
	fmt.Println(string(m1JSON))
}
