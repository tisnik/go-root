package slices2

import "fmt"

func Slices() {
	var a [0]int
	s := a[:]

	for i := 1; i <= 1000000; i++ {
		s = append(s, i)
	}

	fmt.Printf("Length: %d\n", len(s))
}
