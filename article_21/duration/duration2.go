// Seriál "Programovací jazyk Go"
//
// Dvacátá první část
//     Knihovny pro Go umožňující naplánování a spouštění periodických úloh
//     https://www.root.cz/clanky/knihovny-pro-go-umoznujici-naplanovani-a-spousteni-periodickych-uloh/
//
// Demonstrační příklad číslo 2:
//     Práce s objektem typu Duration.

package main

import (
	"fmt"
	"time"
)

func main() {
	d, _ := time.ParseDuration("1h6m10s")

	fmt.Println(d.String())
	fmt.Printf("Hours:   %4.2f\n", d.Hours())
	fmt.Printf("Minutes: %2.0f\n", d.Minutes())
	fmt.Printf("Seconds: %2.0f\n", d.Seconds())
	fmt.Printf("ns:      %d\n", d.Nanoseconds())
}
