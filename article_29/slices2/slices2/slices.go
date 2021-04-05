// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá devátá část
//    Trasování a profilování aplikací naprogramovaných v Go
//    https://www.root.cz/clanky/trasovani-a-profilovani-aplikaci-naprogramovanych-v-go/

package slices2

import "fmt"

func Slices() {
	var a [0]int
	s := a[:]

	for i := 1; i <= 1000000; i++ {
		s = append(s, i)
	}

	fmt.Printf("Length: %d\n", len(s))
}
