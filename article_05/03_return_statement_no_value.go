// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 3:
//    Příkaz return ve funkci bez návratové hodnoty
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/03_return_statement_no_value.html

package main

func f1() {
	println("f1() před příkazem return")
	return
	println("f1() po příkazu return")
}

func main() {
	println("Hello world!")
	f1()
}
