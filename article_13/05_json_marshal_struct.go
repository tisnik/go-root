// Seriál "Programovací jazyk Go"
//
// Třináctá část
//
// Demonstrační příklad číslo 5:
//     Marshalling struktury/záznamu do JSONu

package main

import (
	"encoding/json"
	"fmt"
)

type User1 struct {
	id      uint32
	name    string
	surname string
}

type User2 struct {
	Id      uint32
	Name    string
	Surname string
}

func main() {
	user1 := User1{
		1,
		"Pepek",
		"Vyskoč"}

	user2 := User2{
		1,
		"Pepek",
		"Vyskoč"}

	user1_json, _ := json.Marshal(user1)
	fmt.Println(string(user1_json))

	user2_json, _ := json.Marshal(user2)
	fmt.Println(string(user2_json))
}
