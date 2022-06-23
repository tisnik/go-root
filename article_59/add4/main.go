// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá devátá část
//    Generické datové typy v jazyce Go?
//    https://www.root.cz/clanky/genericke-datove-typy-v-jazyce-go/

package main

import "fmt"
import "add4/adder"

func main() {
	fmt.Println(adder.IntAdd(1, 2))
	fmt.Println(adder.Float32Add(1.1, 2.2))
}
