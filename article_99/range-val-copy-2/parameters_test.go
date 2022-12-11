package main

import (
	"testing"
)

const PAYLOAD_SIZE = 100

type item struct {
	value   int
	payload [PAYLOAD_SIZE]byte
}

const MAX_ITEMS = 10000

func BenchmarkCountValues1(b *testing.B) {
	b.StopTimer()

	var items map[int]item = make(map[int]item, MAX_ITEMS)

	for i := 0; i < MAX_ITEMS; i++ {
		items[i] = item{
			value:   i,
			payload: [PAYLOAD_SIZE]byte{},
		}
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sum := 0

		for _, item := range items {
			sum += item.value
		}

		if sum != MAX_ITEMS/2*(MAX_ITEMS-1) {
			b.Fatal(sum, MAX_ITEMS/2*(MAX_ITEMS-1))
		}
	}
}

func BenchmarkCountValues2(b *testing.B) {
	b.StopTimer()

	var items [MAX_ITEMS]item

	for i := 0; i < len(items); i++ {
		items[i].value = i
		items[i].payload = [PAYLOAD_SIZE]byte{}
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sum := 0

		for key, _ := range items {
			sum += items[key].value
		}

		if sum != MAX_ITEMS/2*(MAX_ITEMS-1) {
			b.Fatal(sum, MAX_ITEMS/2*(MAX_ITEMS-1))
		}
	}
}
