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

func main() {
	var a bool = true

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(a)
	if err != nil {
		fmt.Println(err)
	} else {
		content := buffer.Bytes()
		fmt.Printf("Encoded into %d bytes\n", len(content))
		encoded := hex.EncodeToString(content)
		fmt.Println(encoded)
	}
}
