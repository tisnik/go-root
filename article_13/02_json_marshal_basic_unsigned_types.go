// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 2:
//     Marshalling celých čísel bez znaménka do JSONu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_13/02_json_marshal_basic_unsigned_types.html

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var b8 byte = 0x42
	var a uint8 = 10
	var b uint16 = 1000
	var c uint32 = 10000
	var d uint32 = 1000000

	b8JSON, _ := json.Marshal(b8)
	fmt.Println(string(b8JSON))

	aJSON, _ := json.Marshal(a)
	fmt.Println(string(aJSON))

	bJSON, _ := json.Marshal(b)
	fmt.Println(string(bJSON))

	cJSON, _ := json.Marshal(c)
	fmt.Println(string(cJSON))

	dJSON, _ := json.Marshal(d)
	fmt.Println(string(dJSON))
}
