// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 31:
//    Dvojice vnořených smyček for a příkaz continue.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/31_continue_from_inner_loop.html

package main

import "fmt"

func main() {
Exit:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%3d ", i*j)
			if i*j == 42 {
				fmt.Println("\nodpověď nalezena!\nzkusím další řadu")
				continue Exit
			}
		}
		fmt.Println()
	}
}
