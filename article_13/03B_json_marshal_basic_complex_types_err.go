// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze třinácté části:
//    https://github.com/tisnik/go-root/blob/master/article_13/README.md
//
// Demonstrační příklad číslo 3B:
//     Pokus o marshalling komplexních čísel do JSONu, kontrola převodu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_13/03B_json_marshal_basic_complex_types_err.html

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
