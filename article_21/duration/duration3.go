package main

import (
	"fmt"
	"time"
)

func main() {
	d, _ := time.ParseDuration("200ms")

	fmt.Println(d.String())
	fmt.Printf("Seconds: %4.2f\n", d.Seconds())
	fmt.Printf("ns:      %d\n", d.Nanoseconds())
}
