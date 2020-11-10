// Seriál "Programovací jazyk Go"
//
// Třicátá sedmá část
//    Monitoring služeb a mikroslužeb psaných v Go nástrojem Prometheus
//    https://www.root.cz/clanky/monitoring-sluzeb-a-mikrosluzeb-psanych-v-go-nastrojem-prometheus/

package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"net/http"
	"os"
	"time"
)

var pageRequests = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "page_requests",
	Help: "The total number page/URL requests",
}, []string{"url"})

var histogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "response_time",
	Help:    "Response time",
	Buckets: prometheus.LinearBuckets(0, 5, 20),
}, []string{"url"})

func countEndpoint(request *http.Request, start time.Time) {
	url := request.URL.String()
	fmt.Printf("Request URL: %s\n", url)
	duration := time.Since(start)
	fmt.Printf("Time to serve the page: %s\n", duration)

	// uprava citacu stranek
	pageRequests.With(prometheus.Labels{"url": url}).Inc()

	// uprava histogramu
	histogram.With(prometheus.Labels{"url": url}).Observe(float64(duration.Microseconds()))
}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	io.WriteString(writer, "Hello world!\n")
	countEndpoint(request, start)
}

func fooEndpoint(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	countEndpoint(request, start)
	io.WriteString(writer, "FOO!\n")
}

func barEndpoint(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	io.WriteString(writer, "BAR!\n")
	countEndpoint(request, start)
}

func main() {
	fmt.Println("Initializing HTTP server")

	http.HandleFunc("/", mainEndpoint)
	http.HandleFunc("/foo", fooEndpoint)
	http.HandleFunc("/bar", barEndpoint)
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
