package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var durations = promauto.NewHistogram(prometheus.HistogramOpts{
	Name:    "request_durations",
	Help:    "Durations of requests.",
	Buckets: prometheus.LinearBuckets(0, 0.1, 10),
})

func tick() {
	for {
		v := rand.Float64()
		fmt.Println(v)
		durations.Observe(v)
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
