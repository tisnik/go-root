// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá sedmá část
//    Serializace a deserializace datových struktur v programovacím jazyce Go
//    https://www.root.cz/clanky/serializace-a-deserializace-datovych-struktur-v-programovacim-jazyce-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z padesáté sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_57/README.md

// https://stackoverflow.com/questions/1423081/json-left-out-infinity-and-nan-json-status-in-ecmascript

package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func main() {
	var a float64 = -0.0
	var b float64 = math.NaN()
	var c float64 = -math.NaN()
	var d float64 = math.Inf(1)
	var e float64 = math.Inf(-1)

	var jsonOutput []byte
	var err error

	jsonOutput, err = json.Marshal(a)
	fmt.Println(err, string(jsonOutput))

	jsonOutput, err = json.Marshal(b)
	fmt.Println(err, string(jsonOutput))

	jsonOutput, err = json.Marshal(c)
	fmt.Println(err, string(jsonOutput))

	jsonOutput, err = json.Marshal(d)
	fmt.Println(err, string(jsonOutput))

	jsonOutput, err = json.Marshal(e)
	fmt.Println(err, string(jsonOutput))
}
