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
// Demonstrační příklad číslo 5:
//     Práce s objektem typu Duration.

package main

import (
	"fmt"
	"time"
)

func main() {
	d, _ := time.ParseDuration("3h15m")
	e, _ := time.ParseDuration("1h")
	f := d.Round(e)

	fmt.Println(f.String())
	fmt.Printf("Hours:   %2.0f\n", f.Hours())
	fmt.Printf("Minutes: %2.0f\n", f.Minutes())
	fmt.Printf("Seconds: %4.0f\n", f.Seconds())
	fmt.Printf("ns:      %d\n", f.Nanoseconds())
}
