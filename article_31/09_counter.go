package main

import (
	"fmt"
	"time"
)

func main() {
	var cnt int = 0

	for i := 0; i < 1000; i++ {
		go func() {
			cnt++
		}()
	}
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("%d\n", cnt)
}
