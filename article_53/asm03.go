// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá třetí část
//    Programovací jazyk Go a assembler
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-assembler/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z padesáté třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_53/README.md

package main

func Add(x int, y int) int {
	return x + y
}

func main() {
	println(Add(1, 2))
}
