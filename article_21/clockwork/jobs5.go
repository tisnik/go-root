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
//     Osm naplánovaných úloh vytvořených nástrojem clockwork.

package main

import (
	"github.com/whiteShtef/clockwork"
)

func task1() {
	println("task/job #1 called")
}

func task2() {
	println("task/job #2 called")
}

func task3() {
	println("task/job #2 called")
}

func task4() {
	println("task/job #4 called")
}

func task5() {
	println("task/job #5 called")
}

func task6() {
	println("task/job #6 called")
}

func task7() {
	println("task/job #7 called")
}

func task8() {
	println("task/job #8 called")
}

func main() {
	scheduler := clockwork.NewScheduler()
	scheduler.Schedule().Every(30).Seconds().Do(task1)
	scheduler.Schedule().Every(30).Minutes().Do(task2)
	scheduler.Schedule().Every().Hours().Do(task3)
	scheduler.Schedule().Every(2).Days().Do(task4)
	scheduler.Schedule().Every(2).Days().At("23:59").Do(task5)
	scheduler.Schedule().Every(4).Weeks().Do(task6)
	scheduler.Schedule().Every().Tuesday().Do(task7)
	scheduler.Schedule().Every().Friday().At("19:55").Do(task8)
	scheduler.Run()
}
