// Seriál "Programovací jazyk Go"
//
// Dvacátá první část
//     Knihovny pro Go umožňující naplánování a spouštění periodických úloh
//     https://www.root.cz/clanky/knihovny-pro-go-umoznujici-naplanovani-a-spousteni-periodickych-uloh/
//
// Demonstrační příklad číslo 3:
//     Práce s objektem typu Duration.

package main

import (
	"fmt"
	"time"
)

func main() {
	d, _ := time.ParseDuration("200ms")

	fmt.Println(d.String())
	fmt.Printf("Seconds: %4.2f\n", d.Seconds())
	fmt.Printf("ns:      %d\n", d.Nanoseconds())
}
