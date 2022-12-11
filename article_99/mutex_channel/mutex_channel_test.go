package main

import (
	"sync"
	"testing"
)

type valueWithMutex struct {
	i     int
	mutex *sync.Mutex
}

type valueWithChannel struct {
	i       int
	channel chan struct{}
}

func (value *valueWithMutex) lock() {
	value.mutex.Lock()
}

func (value *valueWithMutex) unlock() {
	value.mutex.Unlock()
}

func (value *valueWithChannel) lock() {
	<-value.channel
}

func (value *valueWithChannel) unlock() {
	value.channel <- struct{}{}
}

const MAX_ITEMS = 100
const EXPECTED_SUM = MAX_ITEMS / 2 * (MAX_ITEMS - 1)

func BenchmarkAccumulationMutexVariant(b *testing.B) {
	var waitgroup sync.WaitGroup
	value := &valueWithMutex{
		mutex: new(sync.Mutex),
		i:     0,
	}

	for n := 0; n < b.N; n++ {
		value.i = 0
		for i := 0; i < MAX_ITEMS; i++ {
			waitgroup.Add(1)
			go func(i int) {
				value.lock()
				defer value.unlock()
				defer waitgroup.Done()
				value.i += i
			}(i)
		}
		waitgroup.Wait()
		if value.i != EXPECTED_SUM {
			b.Fatal(value.i, EXPECTED_SUM)
		}
	}
}

func BenchmarkAccumulationChannelVariant(b *testing.B) {
	var waitgroup sync.WaitGroup
	value := &valueWithChannel{
		channel: make(chan struct{}, 1),
		i:       0,
	}
	value.channel <- struct{}{}

	for n := 0; n < b.N; n++ {
		value.i = 0

		for i := 0; i < MAX_ITEMS; i++ {
			waitgroup.Add(1)
			go func(i int) {
				value.lock()
				defer value.unlock()
				defer waitgroup.Done()
				value.i += i
			}(i)
		}
		waitgroup.Wait()
		if value.i != EXPECTED_SUM {
			b.Fatal(value.i, EXPECTED_SUM)
		}
	}
}
