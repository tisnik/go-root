// Seriál "Programovací jazyk Go"
//
// Devátá část
//
// Demonstrační příklad číslo 8:
//    Základní funkce pro formátování řetězců: FormatInt.

package main

import (
	"strconv"
)

func main() {
	value := int64(0xcafebabe)
	for base := 2; base < 36; base++ {
		println(base, strconv.FormatInt(value, base))
	}
}
