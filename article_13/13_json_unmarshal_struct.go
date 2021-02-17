// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 13:
//     Unmarshalling struktury z JSONu

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// User je uživatelsky definovaná datová struktura s veřejnými atributy
type User struct {
	ID      uint32
	Name    string
	Surname string
}

func main() {
	inputJSONAsBytes, err := ioutil.ReadFile("user.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Input (bytes):")
	fmt.Println(inputJSONAsBytes)

	fmt.Println("\nInput (string):")
	fmt.Println(string(inputJSONAsBytes))

	var user User
	json.Unmarshal(inputJSONAsBytes, &user)

	fmt.Println("\nOutput:")
	fmt.Println(user)

	fmt.Println("\nFields:")
	fmt.Printf("ID:      %d\n", user.ID)
	fmt.Printf("Name:    %s\n", user.Name)
	fmt.Printf("Surname: %s\n", user.Surname)
}
