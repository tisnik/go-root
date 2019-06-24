package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var cnt int32 = 0
	var mutex = &sync.Mutex{}

	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock()
			cnt++
			mutex.Unlock()
		}()
	}
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("%d\n", cnt)
}
