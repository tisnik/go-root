package main

func Swap(x int64, y int64) (int64, int64) {
	return y, x
}

func main() {
	x, y := Swap(1, 2)
	println(x, y)
}
