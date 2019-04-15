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

func task3(counter int) {
	if counter%10 == 0 {
		println("task/job #3 called")
	}
}

func main() {
	task3cnt := 0

	scheduler := clockwork.NewScheduler()
	scheduler.Schedule().Every(20).Seconds().Do(task1)
	scheduler.Schedule().Every(30).Seconds().Do(task2)
	scheduler.Schedule().Every().Seconds().Do(
		func() { task3cnt++; task3(task3cnt) })

	scheduler.Run()
}
