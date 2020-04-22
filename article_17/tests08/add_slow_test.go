// Seriál "Programovací jazyk Go"
//
// Sedmnáctá část
//     Testování aplikací naprogramovaných v jazyce Go
//     https://www.root.cz/clanky/testovani-aplikaci-naprogramovanych-v-jazyce-go/
//
// Demonstrační příklad číslo 8:
//     Implementace jednotkových testů.

// +build slow

package main

import (
	"fmt"
	"math"
	"testing"
)

type AddTest struct {
	x        int32
	y        int32
	expected int32
}

func checkAdd(t *testing.T, testInputs []AddTest) {
	for _, i := range testInputs {
		result := add(i.x, i.y)
		if result != i.expected {
			msg := fmt.Sprintf("%d + %d should be %d, got %d instead",
				i.x, i.y, i.expected, result)
			t.Error(msg)
		}
	}
}

func TestAddMinValues(t *testing.T) {
	var addTestInput = []AddTest{
		{math.MinInt32, 0, math.MinInt32},
		{math.MinInt32, 1, math.MinInt32 + 1},
	}
	checkAdd(t, addTestInput)
}

func TestAddMaxValues(t *testing.T) {
	var addTestInput = []AddTest{
		{math.MaxInt32, 0, math.MaxInt32},
		{math.MaxInt32, 1, math.MinInt32},
		{math.MaxInt32, math.MinInt32, -1},
	}
	checkAdd(t, addTestInput)
}

func TestAddMinMaxValues(t *testing.T) {
	var addTestInput = []AddTest{
		{math.MinInt32, 0, math.MinInt32},
		{math.MinInt32, 1, math.MinInt32 + 1},
		{math.MaxInt32, 0, math.MaxInt32},
		{math.MaxInt32, 1, math.MinInt32},
		{math.MaxInt32, math.MinInt32, -1},
	}
	checkAdd(t, addTestInput)
}
