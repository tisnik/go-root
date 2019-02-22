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

	a_json, _ := json.Marshal(a)
	fmt.Println(string(a_json))

	b_json, _ := json.Marshal(b)
	fmt.Println(string(b_json))

	c_json, _ := json.Marshal(c)
	fmt.Println(string(c_json))

	d_json, _ := json.Marshal(d)
	fmt.Println(string(d_json))
}
