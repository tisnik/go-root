// Seriál "Programovací jazyk Go"
//
// Třináctá část
//
// Demonstrační příklad číslo 15:
//     Unmarshalling pole struktur

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type User struct {
	Id      uint32
	Name    string
	Surname string
}

func main() {
	input_json_as_bytes, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Input (bytes):")
	fmt.Println(input_json_as_bytes)

	fmt.Println("\nInput (string):")
	fmt.Println(string(input_json_as_bytes))

	var users []User
	json.Unmarshal(input_json_as_bytes, &users)

	fmt.Println("\nOutput:")
	fmt.Println(users)

	fmt.Println("\nUsers:")
	for i, user := range users {
		fmt.Printf("%d\t%d\t%s\t%s\n", i, user.Id, user.Name, user.Surname)
	}
}
