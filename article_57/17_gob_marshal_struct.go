// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá sedmá část
//    Serializace a deserializace datových struktur v programovacím jazyce Go
//    https://www.root.cz/clanky/serializace-a-deserializace-datovych-struktur-v-programovacim-jazyce-go/

package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
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

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(user)
	if err != nil {
		fmt.Println(err)
	} else {
		content := buffer.Bytes()
		fmt.Printf("Encoded into %d bytes\n", len(content))
		encoded := hex.EncodeToString(content)
		fmt.Println(encoded)
	}
}
