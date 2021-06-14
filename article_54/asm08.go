// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá čtvrtá část
//    Programovací jazyk Go a assembler (2.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-assembler-2-cast/

package main

func Sign(value int) int {
	if value < 0 {
		return -1
	} else if value > 0 {
		return 1
	} else {
		return 0
	}

}

func main() {
	println(Sign(-100))
	println(Sign(100))
	println(Sign(0))
}
