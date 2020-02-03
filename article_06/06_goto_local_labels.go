// Seriál "Programovací jazyk Go"
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
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
