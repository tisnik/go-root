package main

import (
	"github.com/rk/go-cron"
	"time"
)

func task(t time.Time) {
	println(t.String())
}

func main() {
	c := make(chan bool)

	cron.NewDailyJob(-1, -1, -1, task)

	<-c
}
