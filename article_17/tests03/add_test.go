// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmnáctá část
//     Testování aplikací naprogramovaných v jazyce Go
//     https://www.root.cz/clanky/testovani-aplikaci-naprogramovanych-v-jazyce-go/
//
// Demonstrační příklad číslo 3:
//     Implementace jednotkových testů.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_17/test03/add_test.html

package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(1, 2)
	if result != 3 {
		t.Error("1+2 should be 3, got ", result, "instead")
	}

	result = add(10, 20)
	if result != 30 {
		t.Error("10+20 should be 30, got ", result, "instead")
	}
}

func TestAddZero(t *testing.T) {
	result := add(1, 0)
	if result != 1 {
		t.Error("1+0 should be 1, got ", result, "instead")
	}
}
