// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Desátá část
//     Užitečné balíčky pro každodenní použití Go (2), porovnání výkonnosti Go s céčkem
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-go-2-porovnani-vykonnosti-go-s-ceckem/
//
// Demonstrační příklad číslo 7:
//    Použití balíčku exec, spuštění externí utility.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_10/07_exec.html

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("date", "--date=next Mon")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Next Monday: %s\n", out)
	}
}
