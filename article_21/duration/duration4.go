package main

import (
	"fmt"
	"time"
)

func main() {
	d, _ := time.ParseDuration("0.1µs1ns")

	fmt.Println(d.String())

	fmt.Printf("ns:      %d\n", d.Nanoseconds())
}
