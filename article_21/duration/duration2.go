package main

import (
	"fmt"
	"time"
)

func main() {
	d, _ := time.ParseDuration("1h6m10s")

	fmt.Println(d.String())
	fmt.Printf("Hours:   %4.2f\n", d.Hours())
	fmt.Printf("Minutes: %2.0f\n", d.Minutes())
	fmt.Printf("Seconds: %2.0f\n", d.Seconds())
	fmt.Printf("ns:      %d\n", d.Nanoseconds())
}
