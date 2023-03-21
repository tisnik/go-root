// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z dvacáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_26/README.md
//
// Demonstrační příklad číslo 12B:
//    Nulový řez a funkce append

package main

import "fmt"

func main() {
	var s0 []int
	s1 := []int{}
	s2 := make([]int, 0)
	s3 := make([]int, 10)

	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println()

	s0 = append(s0, 1, 2, 3)
	s1 = append(s1, 1, 2, 3)
	s2 = append(s2, 1, 2, 3)
	s3 = append(s3, 1, 2, 3)

	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
