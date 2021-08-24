// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 8:
//    Použití příkazu defer.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_06/08_defer_func.html

package main

import "fmt"

func main() {
	defer (func() { fmt.Println("Finished") })()

	for i := 10; i >= 0; i-- {
		fmt.Printf("%2d\n", i)
	}
	fmt.Println("Finishing main() function")
}
