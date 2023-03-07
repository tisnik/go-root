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
// Demonstrační příklad číslo 1:
//     Úloha vytvořená nástrojem clockwork.

package main

import (
	"github.com/whiteShtef/clockwork"
)

func task() {
	println("task/job called")
}

func main() {
	scheduler := clockwork.NewScheduler()
	scheduler.Schedule().Every(4).Seconds().Do(task)
	scheduler.Run()
}
