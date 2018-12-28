// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 6B:
//    Špatné použití příkazu goto.

package main

func a() {
	goto FuncB
}

func b() {
FuncB:
}

func main() {
}
