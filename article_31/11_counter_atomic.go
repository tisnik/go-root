package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var cnt int32 = 0

	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt32(&cnt, 1)
		}()
	}
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("%d\n", cnt)
}
