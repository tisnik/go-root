// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá osmá část
//    Generické datové typy v jazyce Go?
//    https://www.root.cz/clanky/genericke-datove-typy-v-jazyce-go/

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add(x float32, y float32) float32 {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1.1, 2.2))
}
