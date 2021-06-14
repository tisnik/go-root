// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá čtvrtá část
//    Programovací jazyk Go a assembler (2.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-assembler-2-cast/

package main

func Add8(x int8, y int8, z int8) int8 {
	return x + y + z
}

func Add16(x int16, y int16, z int16) int16 {
	return x + y + z
}

func Add32(x int32, y int32, z int32) int32 {
	return x + y + z
}

func Add64(x int64, y int64, z int64) int64 {
	return x + y + z
}

func main() {
	println(Add8(1, 2, 3))
	println(Add16(1, 2, 3))
	println(Add32(1, 2, 3))
	println(Add64(1, 2, 3))
}
