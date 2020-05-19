// Seriál "Programovací jazyk Go"
//
// Dvacátá první část
//     Knihovny pro Go umožňující naplánování a spouštění periodických úloh
//     https://www.root.cz/clanky/knihovny-pro-go-umoznujici-naplanovani-a-spousteni-periodickych-uloh/
//
// Demonstrační příklad číslo 1:
//     Úloha vytvořená nástrojem go-cron.

package main

import (
	"github.com/rk/go-cron"
	"time"
)

func task(t time.Time) {
	println(t.String())
}

func main() {
	c := make(chan bool)

	cron.NewDailyJob(-1, -1, -1, task)

	<-c
}
