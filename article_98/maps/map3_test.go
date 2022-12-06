package main

import (
	"testing"
)

type emptyValue struct{}

func BenchmarkInsertIntoPreallocatedMapIntKeyEmptyValue(b *testing.B) {
	m := make(map[int]emptyValue, b.N)

	for i := 0; i < b.N; i++ {
		m[i] = emptyValue{}
	}
}

func BenchmarkInsertIntoEmptyMapEmptyValue(b *testing.B) {
	m := map[int]emptyValue{}

	for i := 0; i < b.N; i++ {
		m[i] = emptyValue{}
	}
}
