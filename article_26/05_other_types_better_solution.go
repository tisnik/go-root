package main

import "fmt"

func main() {
	var p *int
	var s []int
	var m map[string]int
	var c chan int
	var f func()
	var i interface{}

	fmt.Println(p)
	fmt.Println(s)
	fmt.Println(m)
	fmt.Println(c)
	fmt.Println(f)
	fmt.Println(i)
}
