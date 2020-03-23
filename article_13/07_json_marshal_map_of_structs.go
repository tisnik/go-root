// Seriál "Programovací jazyk Go"
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 7:
//     Marshalling map struktur/záznamů do JSONu

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
	m1 := make(map[string]User)

	m1["user-id-1"] = User{
		Id:      1,
		Name:    "Pepek",
		Surname: "Vyskoč"}

	m1["user-id-2"] = User{
		Id:      2,
		Name:    "Josef",
		Surname: "Vyskočil"}

	m1_json, _ := json.Marshal(m1)
	fmt.Println(string(m1_json))
}
