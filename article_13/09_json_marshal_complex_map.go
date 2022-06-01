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
// Demonstrační příklad číslo 9:
//     Marshalling map struktur/záznamů
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_13/09_json_marshal_complex_map.html

package main

import (
	"encoding/json"
	"fmt"
)

// Identifiers je datová struktura reprezentující ID uživatele v systému
type Identifiers struct {
	UID uint32
	GID uint32
}

// User je uživatelsky definovaná datová struktura
type User struct {
	Name    string
	Surname string
	Sign    []byte
	Enabled bool
	Ids     Identifiers
}

func main() {
	mapOfUsers := make(map[string]User)

	mapOfUsers["user-id-1"] = User{
		Ids:     Identifiers{1, 1},
		Name:    "Pepek",
		Surname: "Vyskoč",
		Enabled: true,
		Sign:    []byte{0, 0, 0, 0}}

	mapOfUsers["user-id-2"] = User{
		Ids:     Identifiers{2, 1},
		Name:    "Josef",
		Surname: "Vyskočil",
		Enabled: false,
		Sign:    []byte{42, 10, 0, 255}}

	mapOfUsers["user-id-3"] = User{
		Ids:     Identifiers{3, 1},
		Name:    "Varel",
		Surname: "Frištenský"}

	mapOfUsersJSON, _ := json.Marshal(mapOfUsers)
	fmt.Println(string(mapOfUsersJSON))
}
