package main

import (
	"testing"

	"github.com/daichi-m/go18ds/maps"
	"github.com/daichi-m/go18ds/maps/hashbidimap"
	"github.com/daichi-m/go18ds/maps/hashmap"
	"github.com/daichi-m/go18ds/maps/linkedhashmap"
	"github.com/daichi-m/go18ds/maps/treebidimap"
	"github.com/daichi-m/go18ds/maps/treemap"

	"github.com/daichi-m/go18ds/utils"
)

const max = 1000

func filledMapCheck(b *testing.B, m maps.Map[int, int]) {
	if m.Empty() {
		b.Fail()
	}

	if m.Size() != b.N {
		b.Fail()
	}
}

func BenchmarkHashMapPut(b *testing.B) {
	m := hashmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Put(i, i)
	}

	filledMapCheck(b, m)
}

func BenchmarkLinkedHashMapPut(b *testing.B) {
	m := linkedhashmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Put(i, i)
	}

	filledMapCheck(b, m)
}

func BenchmarkLinkedTreeMapPut(b *testing.B) {
	m := treemap.NewWithIntComparator[int]()
	for i := 0; i < b.N; i++ {
		m.Put(i, i)
	}

	filledMapCheck(b, m)
}

func BenchmarkHashBidiMapPut(b *testing.B) {
	m := hashbidimap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Put(i, i)
	}

	filledMapCheck(b, m)
}

func BenchmarkTreeBidiMapPut(b *testing.B) {
	m := treebidimap.NewWith(utils.NumberComparator[int], utils.NumberComparator[int])
	for i := 0; i < b.N; i++ {
		m.Put(i, i)
	}

	filledMapCheck(b, m)
}

func fillInMap(m maps.Map[int, int], max int) {
	for i := 0; i < max; i++ {
		m.Put(i, i)
	}
}

func getFromMap(b *testing.B, m maps.Map[int, int], max int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < max; j++ {
			val, found := m.Get(j)
			if !found {
				b.Fail()
			}
			if val != j {
				b.Fail()
			}
		}
	}
}

func BenchmarkHashMapGet(b *testing.B) {
	m := hashmap.New[int, int]()

	fillInMap(m, max)
	getFromMap(b, m, max)
}

func BenchmarkLinkedHashMapGet(b *testing.B) {
	m := linkedhashmap.New[int, int]()

	fillInMap(m, max)
	getFromMap(b, m, max)
}

func BenchmarkLinkedTreeMapGet(b *testing.B) {
	m := treemap.NewWithIntComparator[int]()

	fillInMap(m, max)
	getFromMap(b, m, max)
}

func BenchmarkHashBidiMapGet(b *testing.B) {
	m := hashbidimap.New[int, int]()

	fillInMap(m, max)
	getFromMap(b, m, max)
}

func BenchmarkTreeBidiMapGet(b *testing.B) {
	m := treebidimap.NewWith(utils.NumberComparator[int], utils.NumberComparator[int])

	fillInMap(m, max)
	getFromMap(b, m, max)
}
