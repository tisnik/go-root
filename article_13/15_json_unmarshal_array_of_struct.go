// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Seznam příkladů z třinácté části:
//    https://github.com/tisnik/go-root/blob/master/article_13/README.md
//
// Demonstrační příklad číslo 15:
//     Unmarshalling pole struktur
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_13/15_json_unmarshal_array_of_struct.html

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
	inputJSONAsBytes, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Input (bytes):")
	fmt.Println(inputJSONAsBytes)

	fmt.Println("\nInput (string):")
	fmt.Println(string(inputJSONAsBytes))

	var users []User
	json.Unmarshal(inputJSONAsBytes, &users)

	fmt.Println("\nOutput:")
	fmt.Println(users)

	fmt.Println("\nUsers:")
	for i, user := range users {
		fmt.Printf("%d\t%d\t%s\t%s\n", i, user.ID, user.Name, user.Surname)
	}
}
