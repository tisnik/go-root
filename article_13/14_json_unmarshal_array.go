// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 14:
//     Unmarshalling pole z JSONu

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	inputJSONAsBytes, err := ioutil.ReadFile("numbers.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Input (bytes):")
	fmt.Println(inputJSONAsBytes)

	fmt.Println("\nInput (string):")
	fmt.Println(string(inputJSONAsBytes))

	var numbers []int
	json.Unmarshal(inputJSONAsBytes, &numbers)

	fmt.Println("\nOutput:")
	fmt.Println(numbers)

	fmt.Println("\nItems:")
	for i, item := range numbers {
		fmt.Printf("%d\t%d\n", i, item)
	}
}
