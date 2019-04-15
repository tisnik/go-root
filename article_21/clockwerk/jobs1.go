package main

import (
	"github.com/onatm/clockwerk"
	"time"
)

type Task struct{}

func (t Task) Run() {
	println("task/job called")
}

func main() {
	c := make(chan bool)

	var task Task
	scheduler := clockwerk.New()
	scheduler.Every(1 * time.Second).Do(task)
	scheduler.Start()

	<-c
}
