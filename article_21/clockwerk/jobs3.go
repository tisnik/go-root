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
// Demonstrační příklad číslo 3:
//     Vylepšená úloha vytvořená nástrojem clockwerk.

package main

import (
	"github.com/onatm/clockwerk"
	"time"
)

type Task struct{}

func (t Task) Run() {
	println("task/job #1 called", time.Now().String())
}

func main() {
	c := make(chan bool)

	var task Task

	d, _ := time.ParseDuration("2s")

	scheduler := clockwerk.New()
	scheduler.Every(d).Do(task)
	scheduler.Start()

	<-c
}
