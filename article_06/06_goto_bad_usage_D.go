// Seriál "Programovací jazyk Go"
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
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
