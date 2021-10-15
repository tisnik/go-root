// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 6:
//     Marshalling pole struktur/záznamů do JSONu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_13/06_json_mashal_array_of_struct.html

package main

import (
	"encoding/json"
	"fmt"
)

// User je uživatelsky definovaná datová struktura
type User struct {
	ID      uint32
	Name    string
	Surname string
}

func main() {
	var users = [3]User{
		User{
			ID:      1,
			Name:    "Pepek",
			Surname: "Vyskoč"},
		User{
			ID:      2,
			Name:    "Pepek",
			Surname: "Vyskoč"},
		User{
			ID:      3,
			Name:    "Josef",
			Surname: "Vyskočil"},
	}

	usersJSON, _ := json.Marshal(users)
	fmt.Println(string(usersJSON))
}
