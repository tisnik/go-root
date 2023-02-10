// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze třinácté části:
//    https://github.com/tisnik/go-root/blob/master/article_13/README.md
//
// Demonstrační příklad číslo 10:
//     Změna názvů klíčů pri marshallingu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_13/10_json_different_keys.html

package main

import (
	"encoding/json"
	"fmt"
)

// Identifiers je datová struktura reprezentující ID uživatele v systému
type Identifiers struct {
	UID uint32 `json:"user-id"`
	GID uint32 `json:"group-id"`
}

// User je uživatelsky definovaná datová struktura s veřejnými atributy
type User struct {
	Name    string `json:"user-name"`
	Surname string `json:"user-surname"`
	Sign    []byte
	Enabled bool `json:"user-login-enabled"`
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
