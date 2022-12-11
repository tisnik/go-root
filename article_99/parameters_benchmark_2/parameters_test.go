package main

import (
	"testing"
)

const MAX_VALS = 1000
const DEFAULT_VALUE = 0
const FIRST_VALUE = -1
const LAST_VALUE = -2

func changeMe1(values []int) int {
	values[0] = FIRST_VALUE
	values[MAX_VALS-1] = LAST_VALUE
	return values[MAX_VALS/2]
}

func changeMe2(values [MAX_VALS]int) int {
	values[0] = FIRST_VALUE
	values[MAX_VALS-1] = LAST_VALUE
	return values[MAX_VALS/2]
}

func changeMe3(values *[MAX_VALS]int) int {
	values[0] = FIRST_VALUE
	values[MAX_VALS-1] = LAST_VALUE
	return values[MAX_VALS/2]
}

func BenchmarkPassSlice(b *testing.B) {
	var values []int = make([]int, MAX_VALS)

	for i := 0; i < b.N; i++ {
		r := changeMe1(values)
		if r != DEFAULT_VALUE {
			b.Fatal()
		}
	}
	if values[0] != FIRST_VALUE {
		b.Fatal()
	}
	if values[MAX_VALS-1] != LAST_VALUE {
		b.Fatal()
	}
}

func BenchmarkPassArrayByValue(b *testing.B) {
	var values [MAX_VALS]int = [MAX_VALS]int{DEFAULT_VALUE}

	for i := 0; i < b.N; i++ {
		r := changeMe2(values)
		if r != DEFAULT_VALUE {
			b.Fatal()
		}
	}
	if values[0] != DEFAULT_VALUE {
		b.Fatal()
	}
	if values[MAX_VALS-1] != DEFAULT_VALUE {
		b.Fatal()
	}
}

func BenchmarkPassArrayByReference(b *testing.B) {
	var values [MAX_VALS]int = [MAX_VALS]int{DEFAULT_VALUE}

	for i := 0; i < b.N; i++ {
		r := changeMe3(&values)
		if r != DEFAULT_VALUE {
			b.Fatal()
		}
	}
	if values[0] != FIRST_VALUE {
		b.Fatal()
	}
	if values[MAX_VALS-1] != LAST_VALUE {
		b.Fatal()
	}
}
