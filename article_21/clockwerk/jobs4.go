// Seriál "Programovací jazyk Go"
//
// Dvacátá první část
//     Knihovny pro Go umožňující naplánování a spouštění periodických úloh
//     https://www.root.cz/clanky/knihovny-pro-go-umoznujici-naplanovani-a-spousteni-periodickych-uloh/
//
// Demonstrační příklad číslo 4:
//     Vylepšené úlohy vytvořené nástrojem clockwerk.

package main

import (
	"github.com/onatm/clockwerk"
	"time"
)

type Task1 struct{}
type Task2 struct{}

func (t Task1) Run() {
	println("task/job #1 called", time.Now().String())
}

func (t Task2) Run() {
	println("task/job #2 called", time.Now().String())
}

func main() {
	c := make(chan bool)

	var task1 Task1
	var task2 Task2

	duration1, _ := time.ParseDuration("2s")
	duration2, _ := time.ParseDuration("3s")

	scheduler := clockwerk.New()
	scheduler.Every(duration1).Do(task1)
	scheduler.Every(duration2).Do(task2)
	scheduler.Start()

	<-c
}
