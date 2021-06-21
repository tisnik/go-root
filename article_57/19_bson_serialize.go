// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá sedmá část
//    Serializace a deserializace datových struktur v programovacím jazyce Go
//    https://www.root.cz/clanky/serializace-a-deserializace-datovych-struktur-v-programovacim-jazyce-go/

package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
)

type User struct {
	Id      uint32
	Name    string
	Surname string
}

func main() {
	user := User{
		1,
		"Pepek",
		"Vyskoč"}

	var bsonOutput []byte

	bsonOutput, err := bson.Marshal(user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Encoded into %d bytes\n", len(bsonOutput))
		err := ioutil.WriteFile("1.bson", bsonOutput, 0644)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("And stored into file")
		}
	}
}
