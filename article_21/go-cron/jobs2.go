// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá první část
//     Knihovny pro Go umožňující naplánování a spouštění periodických úloh
//     https://www.root.cz/clanky/knihovny-pro-go-umoznujici-naplanovani-a-spousteni-periodickych-uloh/
//
// Demonstrační příklad číslo 2:
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

	cron.NewDailyJob(cron.ANY, cron.ANY, cron.ANY, task)

	<-c
}
