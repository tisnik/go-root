package main

import (
	"encoding/json"
	"fmt"
)

type Identifiers struct {
	UID uint32
	GID uint32
}

type User struct {
	Name    string
	Surname string
	Sign    []byte
	Enabled bool
	Ids     Identifiers
}

func main() {
	mapOfUsers := make(map[string]User)

	mapOfUsers["user-id-1"] = User{
		Ids:     Identifiers{1, 1},
		Name:    "Pepek",
		Surname: "Vyskoč",
		Enabled: true,
		Sign:    []byte{0, 0, 0, 0}}

	mapOfUsers["user-id-2"] = User{
		Ids:     Identifiers{2, 1},
		Name:    "Josef",
		Surname: "Vyskočil",
		Enabled: false,
		Sign:    []byte{42, 10, 0, 255}}

	mapOfUsers["user-id-3"] = User{
		Ids:     Identifiers{3, 1},
		Name:    "Varel",
		Surname: "Frištenský"}

	mapOfUsers_json, _ := json.Marshal(mapOfUsers)
	fmt.Println(string(mapOfUsers_json))
}
