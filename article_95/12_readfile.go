// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z devadesáté páté části:
//    https://github.com/tisnik/go-root/blob/master/article_95/README.md

package main

import (
	"fmt"
	"log"

	narray "github.com/akualab/narray/na64"
)

func main() {
	vector, err := narray.ReadFile("vector.dat")
	if err != nil {
		log.Panic(err)
	}

	matrix, err := narray.ReadFile("matrix.dat")
	if err != nil {
		log.Panic(err)
	}

	cube, err := narray.ReadFile("cube.dat")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(vector)
	fmt.Println(matrix)
	fmt.Println(cube)
}
