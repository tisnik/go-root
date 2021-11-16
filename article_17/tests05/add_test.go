// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmnáctá část
//     Testování aplikací naprogramovaných v jazyce Go
//     https://www.root.cz/clanky/testovani-aplikaci-naprogramovanych-v-jazyce-go/
//
// Demonstrační příklad číslo 5:
//     Implementace jednotkových testů.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_17/test05/add_test.html

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

func TestAdd(t *testing.T) {
	var addTestInput = []AddTest{
		{0, 0, 0},
		{1, 0, 1},
		{2, 0, 2},
		{2, 1, 3},
		{2, -2, 0},
		{math.MaxInt32, 0, math.MaxInt32},
		{math.MaxInt32, 1, math.MinInt32},
		{math.MaxInt32, math.MinInt32, -1},
	}

	for _, i := range addTestInput {
		result := add(i.x, i.y)
		if result != i.expected {
			msg := fmt.Sprintf("%d + %d should be %d, got %d instead",
				i.x, i.y, i.expected, result)
			t.Error(msg)
		}
	}
}
