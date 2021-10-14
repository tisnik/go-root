// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Desátá část
//     Užitečné balíčky pro každodenní použití Go (2), porovnání výkonnosti Go s céčkem
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-go-2-porovnani-vykonnosti-go-s-ceckem/
//
// Demonstrační příklad číslo 11:
//    Získání základních informací o procesu.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_10/11_process_info.html

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("PID = %d\n", os.Getpid())
	fmt.Printf("Parent PID = %d\n", os.Getppid())
	fmt.Printf("Temp. directory = %s\n", os.TempDir())
	cwd, err := os.Getwd()
	if err == nil {
		fmt.Printf("CWD = %s\n", cwd)
	} else {
		fmt.Printf("can not get CWD")
	}
}
