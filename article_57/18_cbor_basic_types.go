// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá sedmá část
//    Serializace a deserializace datových struktur v programovacím jazyce Go
//    https://www.root.cz/clanky/serializace-a-deserializace-datovych-struktur-v-programovacim-jazyce-go/

package main

import (
	"fmt"
	"github.com/fxamacker/cbor/v2"
)

func main() {
	var a bool = true

	var jsonOutput []byte

	cborOutput, _ = cbor.Marshal(a)
	fmt.Println(string(cborOutput))
}
