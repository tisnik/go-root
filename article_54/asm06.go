// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá čtvrtá část
//    Programovací jazyk Go a assembler (2.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-assembler-2-cast/

package main

func AddFloat32(x float32, y float32, z float32) float32 {
	return x + y + z
}

func AddFloat64(x float64, y float64, z float64) float64 {
	return x + y + z
}

func main() {
	println(AddFloat32(1.0, 2.0, 3.0))
	println(AddFloat64(1.0, 2.0, 3.0))
}
