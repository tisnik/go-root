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
	println("task/job #3 called")
}

func main() {
	scheduler := clockwork.NewScheduler()
	scheduler.Schedule().Every(20).Seconds().Do(task1)
	scheduler.Schedule().Every(30).Seconds().Do(task2)
	scheduler.Schedule().Every().Minutes().Do(task3)
	scheduler.Run()
}
