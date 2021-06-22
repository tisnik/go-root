// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá osmá část
//    Generické datové typy v jazyce Go?
//    https://www.root.cz/clanky/genericke-datove-typy-v-jazyce-go/

package main

import "fmt"
import "add3/adder"

func main() {
	fmt.Println(adder.Add(1, 2))
	fmt.Println(adder.Add(1.1, 2.2))
}
