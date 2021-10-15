// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Demonstrační příklad číslo 11:
//     Pokus o marshalling speciálních hodnot
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_13/11_json_marshal_special_values.html

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	v1 := ""
	v2 := false
	v3 := true
	var v4 *int

	var slice1 []int
	slice2 := []int{}

	var map1 map[string]string
	map2 := make(map[string]string)
	map3 := map[string]string{}

	// https://speakerdeck.com/campoy/understanding-nil
	var iface interface{} = nil

	v1JSON, _ := json.Marshal(v1)
	fmt.Println(string(v1JSON))

	v2JSON, _ := json.Marshal(v2)
	fmt.Println(string(v2JSON))

	v3JSON, _ := json.Marshal(v3)
	fmt.Println(string(v3JSON))

	v4JSON, _ := json.Marshal(v4)
	fmt.Println(string(v4JSON))

	slice1JSON, _ := json.Marshal(slice1)
	fmt.Println(string(slice1JSON))

	slice2JSON, _ := json.Marshal(slice2)
	fmt.Println(string(slice2JSON))

	map1JSON, _ := json.Marshal(map1)
	fmt.Println(string(map1JSON))

	map2JSON, _ := json.Marshal(map2)
	fmt.Println(string(map2JSON))

	map3JSON, _ := json.Marshal(map3)
	fmt.Println(string(map3JSON))

	ifaceJSON, _ := json.Marshal(iface)
	fmt.Println(string(ifaceJSON))

	var f = func() {}
	fJSON, _ := json.Marshal(f)
	fmt.Println(string(fJSON))
}
