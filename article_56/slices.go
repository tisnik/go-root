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
