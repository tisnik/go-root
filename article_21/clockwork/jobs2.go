package main

import (
	"github.com/whiteShtef/clockwork"
)

func task() {
	println("task/job called")
}

func main() {
	scheduler := clockwork.NewScheduler()
	scheduler.Schedule().Every().Seconds().Do(task)
	scheduler.Run()
}
