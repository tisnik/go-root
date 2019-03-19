// Seriál "Programovací jazyk Go"
//
// Třináctá část
//
// Demonstrační příklad číslo 7:
//     Marshalling map struktur/záznamů do JSONu

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id      uint32
	Name    string
	Surname string
}

func main() {
	m1 := make(map[string]User)

	m1["user-id-1"] = User{
		Id:      1,
		Name:    "Pepek",
		Surname: "Vyskoč"}

	m1["user-id-2"] = User{
		Id:      2,
		Name:    "Josef",
		Surname: "Vyskočil"}

	m1_json, _ := json.Marshal(m1)
	fmt.Println(string(m1_json))
}
