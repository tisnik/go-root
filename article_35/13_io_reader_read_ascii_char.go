// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá pátá část
//    Programovací jazyk Go pro skalní céčkaře (2.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const filename = "test_input.txt"

func main() {
	fin, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	buffer := make([]byte, 1)

	for {
		read, err := fin.Read(buffer)

		if read > 0 {
			fmt.Printf("%c", buffer[0])
		}

		if err == io.EOF {
			fmt.Println("\nreached end of file")
			break
		}

		if err != nil {
			fmt.Printf("\nother error %v\n", err)
			break
		}
	}
}
