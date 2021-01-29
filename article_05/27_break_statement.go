// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
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
