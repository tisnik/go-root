// Seriál "Programovací jazyk Go"
//
// Třináctá část
//
// Demonstrační příklad číslo 12:
//     Unmarshalling struktury z JSONu

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
	input_json := `{
    "Id":1,
    "Name":"Pepek",
    "Surname":"Vyskoč"
}`
	fmt.Println("Input:")
	fmt.Println(input_json)

	bytes := []byte(input_json)
	var user User
	json.Unmarshal(bytes, &user)

	fmt.Println("\nOutput:")
	fmt.Println(user)

	fmt.Println("\nFields:")
	fmt.Printf("ID:      %d\n", user.Id)
	fmt.Printf("Name:    %s\n", user.Name)
	fmt.Printf("Surname: %s\n", user.Surname)
}
