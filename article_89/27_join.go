// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z osmdesáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_89/README.md

package main

import "fmt"

func join[T any](items []T) (result string) {
	for _, value := range items {
		result += value.String()
		result += ","
	}
	return result
}

func main() {
	fmt.Println(join([]string{"first", "second"}))
}
