// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 18:
//    Praktické použití konstrukce defer, panic a recover.

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
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("Recovered in copyFile", rec)
		}
	}()
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

	return io.Copy(dst, src)
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
	testCopyFile("14_defer_practical_usage.go", "new.go")
	// testCopyFile("tento_soubor_neexistuje", "new.go")
	testCopyFile("new.go", "")
	testCopyFile("14_defer_practical_usage.go", "/dev/full")
	testCopyFile("/dev/null", "new2.go")
}
