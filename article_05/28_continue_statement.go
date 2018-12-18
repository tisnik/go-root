// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Demonstrační příklad číslo 28:
//    Příkaz continue.

package main

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		println(i)
	}
}
