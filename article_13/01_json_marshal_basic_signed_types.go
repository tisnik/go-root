// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 1:
//     Marshalling celých čísel se znaménkem do JSONu

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a int8 = -10
	var b int16 = -1000
	var c int32 = -10000
	var d int32 = -1000000

	var r1 rune = 'a'
	var r2 rune = '\x40'
	var r3 rune = '\n'
	var r4 rune = '\u03BB'

	aJSON, _ := json.Marshal(a)
	fmt.Println(string(aJSON))

	bJSON, _ := json.Marshal(b)
	fmt.Println(string(bJSON))

	cJSON, _ := json.Marshal(c)
	fmt.Println(string(cJSON))

	dJSON, _ := json.Marshal(d)
	fmt.Println(string(dJSON))

	r1JSON, _ := json.Marshal(r1)
	fmt.Println(string(r1JSON))

	r2JSON, _ := json.Marshal(r2)
	fmt.Println(string(r2JSON))

	r3JSON, _ := json.Marshal(r3)
	fmt.Println(string(r3JSON))

	r4JSON, _ := json.Marshal(r4)
	fmt.Println(string(r4JSON))
}
