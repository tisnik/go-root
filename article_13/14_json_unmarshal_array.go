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
// Demonstrační příklad číslo 14:
//     Unmarshalling pole z JSONu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_13/14_json_unmarshal_array.html

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
