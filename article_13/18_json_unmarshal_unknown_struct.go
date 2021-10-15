// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 18:
//     Pokus o unmarshalling struktury s obecnými daty
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_13/18_json_unmarshal_unknown_struct.html

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	inputJSONAsBytes, err := ioutil.ReadFile("users_map_different_keys.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Input (bytes):")
	fmt.Println(inputJSONAsBytes)

	fmt.Println("\nInput (string):")
	fmt.Println(string(inputJSONAsBytes))

	m1 := map[string]interface{}{}
	json.Unmarshal(inputJSONAsBytes, &m1)

	fmt.Println("\nOutput:")
	fmt.Println(m1)

	fmt.Println("\nUsers:")
	for key, user := range m1 {
		fmt.Printf("%s\t%s\n", key, user)
	}
}
