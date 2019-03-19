// Seriál "Programovací jazyk Go"
//
// Třináctá část
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

	a_json, a_err := json.Marshal(a)
	fmt.Println(string(a_json))
	fmt.Println(a_err)

	b_json, b_err := json.Marshal(b)
	fmt.Println(string(b_json))
	fmt.Println(b_err)

	c_json, c_err := json.Marshal(c)
	fmt.Println(string(c_json))
	fmt.Println(c_err)

	d_json, d_err := json.Marshal(d)
	fmt.Println(string(d_json))
	fmt.Println(d_err)
}
