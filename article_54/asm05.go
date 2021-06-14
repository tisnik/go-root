// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá čtvrtá část
//    Programovací jazyk Go a assembler (2.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-assembler-2-cast/

package main

func Swap(x int64, y int64) (int64, int64) {
	return y, x
}

func main() {
	x, y := Swap(1, 2)
	println(x, y)
}
