// Seriál "Programovací jazyk Go"
//
// Dvacátá druhá část
//
// Demonstrační příklad číslo 18:
//     Blokové přenosy dat při kopii souboru.

package main

import (
	"fmt"
	"io"
	"os"
)

func closeFile(file *os.File) {
	fmt.Printf("Closing file '%s'\n", file.Name())
	file.Close()
}

func copyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		panic(err)
	}
	defer closeFile(src)

	dst, err := os.Create(dstName)
	if err != nil {
		panic(err)
	}
	defer closeFile(dst)

	buffer := make([]byte, 16)
	copied := int64(0)

	for {
		read, err := src.Read(buffer)

		if read > 0 {
			fmt.Printf("read %d bytes\n", read)
			written, err := dst.Write(buffer[:read])
			if written > 0 {
				fmt.Printf("written %d bytes\n", written)
			}
			if err != nil {
				fmt.Printf("write error %v\n", err)
				return copied, err
			}
			copied += int64(written)
		}

		if err == io.EOF {
			fmt.Println("reached end of file")
			break
		}

		if err != nil {
			fmt.Printf("other error %v\n", err)
			return copied, err
		}
	}
	return copied, nil
}

func testCopyFile(srcName, dstName string) {
	copied, err := copyFile(srcName, dstName)
	if err != nil {
		fmt.Printf("copyFile('%s', '%s') failed!!!\n", srcName, dstName)
	} else {
		fmt.Printf("Copied %d bytes\n", copied)
	}
	fmt.Println()
}

func main() {
	testCopyFile("test_input.txt", "output.txt")
}
