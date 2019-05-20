package main

import "fmt"

func main() {
	var a [10]complex64
	var p *int
	var s []int
	var m map[string]int
	var c chan int
	var f func()
	var i interface{}

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(s)
	fmt.Println(m)
	fmt.Println(c)
	fmt.Println(f)
	fmt.Println(i)
}
