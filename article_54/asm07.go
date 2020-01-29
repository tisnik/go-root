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
