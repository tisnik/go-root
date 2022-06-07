// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Seznam příkladů ze třinácté části:
//    https://github.com/tisnik/go-root/blob/master/article_13/README.md
//
// Demonstrační příklad číslo 6B:
//     Marshalling pole struktur/záznamů do JSONu, privátní prvky struktury
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_13/06B_json_mashal_array_of_struct.html

package main

import (
	"encoding/json"
	"fmt"
)

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	var users = [3]User{
		User{
			id:      1,
			name:    "Pepek",
			surname: "Vyskoč"},
		User{
			id:      2,
			name:    "Pepek",
			surname: "Vyskoč"},
		User{
			id:      3,
			name:    "Josef",
			surname: "Vyskočil"},
	}

	usersJSON, _ := json.Marshal(users)
	fmt.Println(string(usersJSON))
}
