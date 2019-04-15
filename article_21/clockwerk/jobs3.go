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
