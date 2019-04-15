package main

import (
	"fmt"
	"time"
)

func main() {
	d, _ := time.ParseDuration("1h")

	fmt.Println(d.String())
	fmt.Printf("Hours:   %2.0f\n", d.Hours())
	fmt.Printf("Minutes: %2.0f\n", d.Minutes())
}
