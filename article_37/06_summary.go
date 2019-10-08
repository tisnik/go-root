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

var durations = promauto.NewSummary(prometheus.SummaryOpts{
	Name:       "request_durations",
	Help:       "Durations of requests.",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
})

func tick() {
	for {
		v := rand.Float64()
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
