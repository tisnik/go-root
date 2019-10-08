package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"time"
)

var counter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "number_of_ticks",
	Help: "The total number of ticks since the app is started",
})

func tick() {
	for {
		counter.Inc()
		time.Sleep(time.Second)
	}
}

func recordMetrics() {
	fmt.Println("Starting recording metrics")
	go tick()
}

func main() {
	recordMetrics()
	fmt.Println("Initializing HTTP server")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
