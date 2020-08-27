// Seriál "Programovací jazyk Go"
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 12:
//     Unmarshalling struktury z JSONu

package main

import (
	"encoding/json"
	"fmt"
)

// User je uživatelsky definovaná datová struktura s veřejnými atributy
type User struct {
	Id      uint32
	Name    string
	Surname string
}

func main() {
	input_json := `{
    "Id":1,
    "Name":"Pepek",
    "Surname":"Vyskoč"
}`
	fmt.Println("Input:")
	fmt.Println(input_json)

	bytes := []byte(input_json)
	var user User
	json.Unmarshal(bytes, &user)

	fmt.Println("\nOutput:")
	fmt.Println(user)

	fmt.Println("\nFields:")
	fmt.Printf("ID:      %d\n", user.Id)
	fmt.Printf("Name:    %s\n", user.Name)
	fmt.Printf("Surname: %s\n", user.Surname)
}
