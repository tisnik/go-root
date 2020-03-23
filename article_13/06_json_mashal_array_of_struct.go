// Seriál "Programovací jazyk Go"
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 6:
//     Marshalling pole struktur/záznamů do JSONu

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id      uint32
	Name    string
	Surname string
}

func main() {
	var users = [3]User{
		User{
			Id:      1,
			Name:    "Pepek",
			Surname: "Vyskoč"},
		User{
			Id:      2,
			Name:    "Pepek",
			Surname: "Vyskoč"},
		User{
			Id:      3,
			Name:    "Josef",
			Surname: "Vyskočil"},
	}

	users_json, _ := json.Marshal(users)
	fmt.Println(string(users_json))
}
