// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Seznam příkladů ze šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_06/README.md
//
// Demonstrační příklad číslo 7:
//    Použití příkazu defer.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_06/07_defer.html

package main

import "fmt"

func onFinish() {
	fmt.Println("Finished")
}

func main() {
	defer onFinish()

	for i := 10; i >= 0; i-- {
		fmt.Printf("%2d\n", i)
	}
	fmt.Println("Finishing main() function")
}
