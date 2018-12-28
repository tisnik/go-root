// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 6C:
//    Lokální návěští.

package main

func a() {
	goto Label
Label:
}

func b() {
	goto Label
Label:
}

func main() {
}
