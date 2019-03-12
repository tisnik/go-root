// Seriál "Programovací jazyk Go"
//
// Třináctá část
//
// Demonstrační příklad číslo 1:
//    

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

	a_json, _ := json.Marshal(a)
	fmt.Println(string(a_json))

	b_json, _ := json.Marshal(b)
	fmt.Println(string(b_json))

	c_json, _ := json.Marshal(c)
	fmt.Println(string(c_json))

	d_json, _ := json.Marshal(d)
	fmt.Println(string(d_json))

	r1_json, _ := json.Marshal(r1)
	fmt.Println(string(r1_json))

	r2_json, _ := json.Marshal(r2)
	fmt.Println(string(r2_json))

	r3_json, _ := json.Marshal(r3)
	fmt.Println(string(r3_json))

	r4_json, _ := json.Marshal(r4)
	fmt.Println(string(r4_json))
}
