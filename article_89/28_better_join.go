// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/

package main

import "fmt"

func join[T any](items ...T) (result string) {
	for _, value := range items {
		result += value.String()
		result += ","
	}
	return result
}

func main() {
	fmt.Println(join("first", "second"))
}
