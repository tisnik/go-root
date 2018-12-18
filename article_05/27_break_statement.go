// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Demonstrační příklad číslo 27:
//    Příkaz break.

package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
		if i == 5 {
			break
		}
	}
}
