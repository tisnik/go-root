// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 12:
//     Unmarshalling struktury z JSONu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_13/12_json_unmarshal_struct.html

package main

import (
	"encoding/json"
	"fmt"
)

// User je uživatelsky definovaná datová struktura s veřejnými atributy
type User struct {
	ID      uint32
	Name    string
	Surname string
}

func main() {
	inputJSON := `{
    "ID":1,
    "Name":"Pepek",
    "Surname":"Vyskoč"
}`
	fmt.Println("Input:")
	fmt.Println(inputJSON)

	bytes := []byte(inputJSON)
	var user User
	json.Unmarshal(bytes, &user)

	fmt.Println("\nOutput:")
	fmt.Println(user)

	fmt.Println("\nFields:")
	fmt.Printf("ID:      %d\n", user.ID)
	fmt.Printf("Name:    %s\n", user.Name)
	fmt.Printf("Surname: %s\n", user.Surname)
}
