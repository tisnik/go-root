// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá devátá část
//    Generické datové typy v jazyce Go?
//    https://www.root.cz/clanky/genericke-datove-typy-v-jazyce-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z padesáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_59/README.md

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
}
