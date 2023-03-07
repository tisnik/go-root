// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá první část
//     Knihovny pro Go umožňující naplánování a spouštění periodických úloh
//     https://www.root.cz/clanky/knihovny-pro-go-umoznujici-naplanovani-a-spousteni-periodickych-uloh/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z dvacáté první části:
//    https://github.com/tisnik/go-root/blob/master/article_21/README.md
//
// Demonstrační příklad číslo 4:
//     Práce s objektem typu Duration.

package main

import (
	"fmt"
	"time"
)

func main() {
	d, _ := time.ParseDuration("0.1µs1ns")

	fmt.Println(d.String())

	fmt.Printf("ns:      %d\n", d.Nanoseconds())
}
