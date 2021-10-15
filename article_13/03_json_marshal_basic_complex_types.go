// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 3:
//     Pokus o marshalling komplexních čísel do JSONu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_13/03_json_marshal_basic_complex_types.html

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

	aJSON, _ := json.Marshal(a)
	fmt.Println(string(aJSON))

	bJSON, _ := json.Marshal(b)
	fmt.Println(string(bJSON))

	cJSON, _ := json.Marshal(c)
	fmt.Println(string(cJSON))

	dJSON, _ := json.Marshal(d)
	fmt.Println(string(dJSON))
}
