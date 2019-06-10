package slices

import "fmt"

func Slices() {
	var a [0]int
	s := a[:]

	fmt.Println(s)

	for i := 1; i <= 100000; i++ {
		s = append(s, i)
	}

	fmt.Println(s)
}
