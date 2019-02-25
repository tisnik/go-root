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
	input_json_as_bytes, err := ioutil.ReadFile("user.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Input (bytes):")
	fmt.Println(input_json_as_bytes)

	fmt.Println("\nInput (string):")
	fmt.Println(string(input_json_as_bytes))

	var user User
	json.Unmarshal(input_json_as_bytes, &user)

	fmt.Println("\nOutput:")
	fmt.Println(user)

	fmt.Println("\nFields:")
	fmt.Printf("ID:      %d\n", user.Id)
	fmt.Printf("Name:    %s\n", user.Name)
	fmt.Printf("Surname: %s\n", user.Surname)
}
