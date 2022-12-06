package main

import (
	"testing"
)

func BenchmarkInsertIntoPreallocatedMapIntKeyIntValue(b *testing.B) {
	m := make(map[int]int, b.N)

	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}

func BenchmarkInsertIntoEmptyMapIntKeyIntValue(b *testing.B) {
	m := map[int]int{}

	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}
