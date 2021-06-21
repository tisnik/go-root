// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá šestá část
//    Programovací jazyk Go a assembler (3.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-assembler-3-cast/

package main

func GetLen(b []byte) int {
	return len(b)
}

func GetCap(b []byte) int {
	return cap(b)
}

func main() {
	var x []byte = []byte{1, 2, 3}
	println(GetLen(x))
	println(GetCap(x))
}
