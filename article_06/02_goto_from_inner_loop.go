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
// Demonstrační příklad číslo 2:
//    Dvojice vnořených smyček for a příkaz goto pro jejich opuštění.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_06/02_goto_from_inner_loop.html

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%3d ", i*j)
			if i*j == 42 {
				fmt.Println("\nodpověď nalezena!\n")
				goto Exit
			}
		}
		fmt.Println()
	}
Exit:
}
