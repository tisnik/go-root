// Seriál "Programovací jazyk Go"
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 16:
//     Unmarshalling mapy struktur

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
	inputJSONAsBytes, err := ioutil.ReadFile("users_map.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Input (bytes):")
	fmt.Println(inputJSONAsBytes)

	fmt.Println("\nInput (string):")
	fmt.Println(string(inputJSONAsBytes))

	m1 := map[string]User{}
	json.Unmarshal(inputJSONAsBytes, &m1)

	fmt.Println("\nOutput:")
	fmt.Println(m1)

	fmt.Println("\nUsers:")
	for key, user := range m1 {
		fmt.Printf("%s\t%d\t%s\t%s\n", key, user.ID, user.Name, user.Surname)
	}
}
