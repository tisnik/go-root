// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá první část
//     Knihovny pro Go umožňující naplánování a spouštění periodických úloh
//     https://www.root.cz/clanky/knihovny-pro-go-umoznujici-naplanovani-a-spousteni-periodickych-uloh/
//
// Seznam příkladů z dvacáté první části:
//    https://github.com/tisnik/go-root/blob/master/article_21/README.md
//
// Demonstrační příklad číslo 4:
//     Tři úlohy vytvořené nástrojem go-cron.

package main

import (
	"github.com/rk/go-cron"
	"time"
)

func task1(t time.Time) {
	println("task1:", t.String())
}

func task2(t time.Time) {
	println("task2:", t.String())
}

func task3(t time.Time, counter int) {
	if counter%10 == 0 {
		println("task3:", t.String())
	}
}

func main() {
	task3cnt := 0

	c := make(chan bool)
	cron.NewDailyJob(cron.ANY, cron.ANY, 0, task1)
	cron.NewDailyJob(cron.ANY, cron.ANY, 10, task2)
	cron.NewDailyJob(cron.ANY, cron.ANY, cron.ANY,
		func(t time.Time) { task3cnt++; task3(t, task3cnt) })

	<-c
}
