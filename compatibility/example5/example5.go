package main

import "fmt"

func main() {

	var f func()
	for i := 0; i < 10; i++ {
		if i == 0 {
			f = func() { fmt.Print(i) }
		}
		f()
	}
}
