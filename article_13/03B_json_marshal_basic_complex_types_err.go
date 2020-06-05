// Seriál "Programovací jazyk Go"
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 3B:
//     Pokus o marshalling komplexních čísel do JSONu, kontrola převodu

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a complex64 = -1.5 + 0i
	var b complex64 = 1.5 + 1000i
	var c complex64 = 1e30 + 1e30i
	var d complex64 = 1i

	aJSON, aErr := json.Marshal(a)
	fmt.Println(string(aJSON))
	fmt.Println(aErr)

	bJSON, bErr := json.Marshal(b)
	fmt.Println(string(bJSON))
	fmt.Println(bErr)

	cJSON, cErr := json.Marshal(c)
	fmt.Println(string(cJSON))
	fmt.Println(cErr)

	dJSON, dErr := json.Marshal(d)
	fmt.Println(string(dJSON))
	fmt.Println(dErr)
}
