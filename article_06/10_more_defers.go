// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_06/README.md
//
// Demonstrační příklad číslo 10:
//    Zjištění pořadí volání funkcí deklarovaných v defer().
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_06/10_more_defers.html

package main

import "fmt"

func onFinish(i int) {
	fmt.Printf("Defer #%2d\n", i)
}

func main() {
	for i := 0; i <= 10; i++ {
		defer onFinish(i)
	}
	fmt.Println("Finishing main() function")
}
