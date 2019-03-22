// Seriál "Programovací jazyk Go"
//
// Třináctá část
//
// Demonstrační příklad číslo 16:
//     Unmarshalling mapy struktur

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
	input_json_as_bytes, err := ioutil.ReadFile("users_map.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Input (bytes):")
	fmt.Println(input_json_as_bytes)

	fmt.Println("\nInput (string):")
	fmt.Println(string(input_json_as_bytes))

	m1 := map[string]User{}
	json.Unmarshal(input_json_as_bytes, &m1)

	fmt.Println("\nOutput:")
	fmt.Println(m1)

	fmt.Println("\nUsers:")
	for key, user := range m1 {
		fmt.Printf("%s\t%d\t%s\t%s\n", key, user.Id, user.Name, user.Surname)
	}
}
