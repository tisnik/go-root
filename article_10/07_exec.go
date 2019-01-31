// Seriál "Programovací jazyk Go"
//
// Desátá část
//
// Demonstrační příklad číslo 7:
//    Použití balíčku exec, spuštění externí utility.

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
