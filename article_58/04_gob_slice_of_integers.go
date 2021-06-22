// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá osmá část
//    Serializace a deserializace datových struktur v programovacím jazyce Go (2.část)
//    https://www.root.cz/clanky/serializace-a-deserializace-datovych-struktur-v-programovacim-jazyce-go-2-cast/

package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
)

func encodeAndDecodeSliceOfInts(values []int64) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(values)
	if err != nil {
		fmt.Println(err)
	} else {
		content := buffer.Bytes()
		fmt.Printf("Slice with %d values encoded into %d bytes: ", len(values), len(content))
		encoded := hex.EncodeToString(content)
		fmt.Println(encoded)
	}
}

func main() {
	values1 := []int64{}
	encodeAndDecodeSliceOfInts(values1)

	values2 := []int64{1, 2, 3, 4}
	encodeAndDecodeSliceOfInts(values2)

	values3 := []int64{1, 2, 3, 4, 5}
	encodeAndDecodeSliceOfInts(values3)

	values4 := []int64{1000000, 2000000, 3000000, 4000000}
	encodeAndDecodeSliceOfInts(values4)
}
