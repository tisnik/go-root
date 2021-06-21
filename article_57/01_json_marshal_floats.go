// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá sedmá část
//    Serializace a deserializace datových struktur v programovacím jazyce Go
//    https://www.root.cz/clanky/serializace-a-deserializace-datovych-struktur-v-programovacim-jazyce-go/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a float64 = 0.0
	var b float64 = 1e10
	var c float64 = 1e100
	var d float64 = 2.3e-50

	var jsonOutput []byte

	jsonOutput, _ = json.Marshal(a)
	fmt.Println(string(jsonOutput))

	jsonOutput, _ = json.Marshal(b)
	fmt.Println(string(jsonOutput))

	jsonOutput, _ = json.Marshal(c)
	fmt.Println(string(jsonOutput))

	jsonOutput, _ = json.Marshal(d)
	fmt.Println(string(jsonOutput))
}
