// Seriál "Programovací jazyk Go"
//
// Dvacátá devátá část
//    Trasování a profilování aplikací naprogramovaných v Go
//    https://www.root.cz/clanky/trasovani-a-profilovani-aplikaci-naprogramovanych-v-go/

package slices

import "fmt"

func Slices() {
	var a [0]int
	s := a[:]

	fmt.Println(s)

	for i := 1; i <= 100000; i++ {
		s = append(s, i)
	}

	fmt.Println(s)
}
