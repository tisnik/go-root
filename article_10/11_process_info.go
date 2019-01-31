// Seriál "Programovací jazyk Go"
//
// Desátá část
//
// Demonstrační příklad číslo 11:
//    Získání základních informací o procesu.

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
