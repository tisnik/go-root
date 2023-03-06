// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmnáctá část
//     Testování aplikací naprogramovaných v jazyce Go
//     https://www.root.cz/clanky/testovani-aplikaci-naprogramovanych-v-jazyce-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze sedmnácté části:
//    https://github.com/tisnik/go-root/blob/master/article_17/README.md
//
// Demonstrační příklad číslo 8:
//     Implementace jednotkových testů.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_17/test08/add_fast_test.html

//go:build fast
// +build fast

package main

import (
	"fmt"
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

func TestAddBasicValues(t *testing.T) {
	var addTestInput = []AddTest{
		{0, 0, 0},
		{1, 0, 1},
		{2, 0, 2},
		{2, 1, 3},
	}
	checkAdd(t, addTestInput)
}

func TestAddNegativeValues(t *testing.T) {
	var addTestInput = []AddTest{
		{0, 0, 0},
		{1, 0, 1},
		{2, 0, 2},
		{2, 1, 3},
		{2, -2, 0},
	}
	checkAdd(t, addTestInput)
}
