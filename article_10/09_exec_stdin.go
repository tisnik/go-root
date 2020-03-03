// Seriál "Programovací jazyk Go"
//
// Desátá část
//     Užitečné balíčky pro každodenní použití Go (2), porovnání výkonnosti Go s céčkem
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-go-2-porovnani-vykonnosti-go-s-ceckem/
//
// Demonstrační příklad číslo 9:
//    Použití balíčku exec, spuštění externí utility s předáním dat.

package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	cmd := exec.Command("sort")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "zzz\n")
		io.WriteString(stdin, "xyz\n")
		io.WriteString(stdin, "aaa\n")
	}()

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("sorted input:\n%s\n", out)
	}
}
