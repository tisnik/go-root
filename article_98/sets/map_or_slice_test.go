package main

import (
	"testing"
)

type ID int

func genID(i int) ID {
	return ID(3*i + 1)
}

func fillInMap(b *testing.B, items int) map[ID]struct{} {
	b.StopTimer()

	m := make(map[ID]struct{}, items)

	for i := 0; i < items; i++ {
		id := genID(i)
		m[id] = struct{}{}
	}

	b.StartTimer()
	return m
}

func fillInSlice(b *testing.B, items int) []ID {
	b.StopTimer()

	s := make([]ID, items)

	for i := 0; i < items; i++ {
		id := genID(i)
		s[i] = id
	}

	b.StartTimer()
	return s
}

func performBenchmarkFindInMap(b *testing.B, m map[ID]struct{}) {
	items := len(m)
	for i := 0; i < b.N; i++ {
		_, found := m[genID(i%items)]
		if !found {
			b.Fatal("not found")
		}
	}
}

func performBenchmarkFindInSlice(b *testing.B, s []ID) {
	items := len(s)
	for i := 0; i < b.N; i++ {
		found := false
		id := genID(i % items)
		for _, p := range s {
			if p == id {
				found = true
				break
			}
		}
		if !found {
			b.Fatal("not found")
		}
	}
}

func BenchmarkFindInMap1(b *testing.B) {
	m := fillInMap(b, 1)

	performBenchmarkFindInMap(b, m)
}

func BenchmarkFindInSlice1(b *testing.B) {
	m := fillInSlice(b, 1)
	performBenchmarkFindInSlice(b, m)
}

func BenchmarkFindInMap5(b *testing.B) {
	m := fillInMap(b, 5)

	performBenchmarkFindInMap(b, m)
}

func BenchmarkFindInSlice5(b *testing.B) {
	m := fillInSlice(b, 5)
	performBenchmarkFindInSlice(b, m)
}

func BenchmarkFindInMap10(b *testing.B) {
	m := fillInMap(b, 10)

	performBenchmarkFindInMap(b, m)
}

func BenchmarkFindInSlice10(b *testing.B) {
	m := fillInSlice(b, 10)
	performBenchmarkFindInSlice(b, m)
}

func BenchmarkFindInMap20(b *testing.B) {
	m := fillInMap(b, 20)

	performBenchmarkFindInMap(b, m)
}

func BenchmarkFindInSlice20(b *testing.B) {
	m := fillInSlice(b, 20)
	performBenchmarkFindInSlice(b, m)
}

func BenchmarkFindInMap100(b *testing.B) {
	m := fillInMap(b, 100)

	performBenchmarkFindInMap(b, m)
}

func BenchmarkFindInSlice100(b *testing.B) {
	m := fillInSlice(b, 100)
	performBenchmarkFindInSlice(b, m)
}

func BenchmarkFindInMap1000(b *testing.B) {
	m := fillInMap(b, 1000)

	performBenchmarkFindInMap(b, m)
}

func BenchmarkFindInSlice1000(b *testing.B) {
	m := fillInSlice(b, 1000)
	performBenchmarkFindInSlice(b, m)
}
