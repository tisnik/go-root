// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá čtvrtá část
//    Programovací jazyk Go a assembler (2.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-assembler-2-cast/

package main

func Sum(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func main() {
	println(Sum([]int{}))
	println(Sum([]int{1, 2, 3, 4, 5}))
}
