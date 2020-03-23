// Seriál "Programovací jazyk Go"
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 10:
//     Změna názvů klíčů pri marshallingu

package main

import (
	"encoding/json"
	"fmt"
)

type Identifiers struct {
	UID uint32 `json:"user-id"`
	GID uint32 `json:"group-id"`
}

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

	mapOfUsers_json, _ := json.Marshal(mapOfUsers)
	fmt.Println(string(mapOfUsers_json))
}
