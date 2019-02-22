package main

import (
	"encoding/json"
	"fmt"
)

type Key struct {
	Id   uint32
	Role string
}

type User struct {
	Id      uint32
	Name    string
	Surname string
}

func main() {
	m1 := make(map[Key]User)

	m1[Key{1, "admin"}] = User{
		Id:      1,
		Name:    "Pepek",
		Surname: "Vyskoč"}

	m1[Key{2, "user"}] = User{
		Id:      2,
		Name:    "Josef",
		Surname: "Vyskočil"}

	m1_json, _ := json.Marshal(m1)
	fmt.Println(string(m1_json))
}
