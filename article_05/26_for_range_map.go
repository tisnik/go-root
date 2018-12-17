// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Demonstrační příklad číslo 26:
//    For s range prováděným nad mapou.

package main

func main() {
	var m1 map[int]string = make(map[int]string)
	m1[0] = "nula"
	m1[1] = "jedna"
	m1[2] = "dva"
	m1[3] = "tri"
	m1[4] = "ctyri"
	m1[5] = "pet"
	m1[6] = "sest"

	for key, val := range m1 {
		println(key, val)
	}
}
