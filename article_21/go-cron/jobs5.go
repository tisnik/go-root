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

func task3(t time.Time) {
	println("task3:", t.String())
}

func main() {
	c := make(chan bool)

	cron.NewWeeklyJob(cron.ANY, 21, 05, 00, task1)
	cron.NewMonthlyJob(cron.ANY, 21, 05, 00, task2)
	cron.NewCronJob(cron.ANY, cron.ANY, cron.ANY, 21, 05, 00, task3)

	<-c
}
