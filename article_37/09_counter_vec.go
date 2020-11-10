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
)

var pageRequests = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "page_requests",
	Help: "The total number page/URL requests",
},
	[]string{"url"})

func countEndpoint(request *http.Request) {
	url := request.URL.String()
	fmt.Printf("Request URL: %s\n", url)
	pageRequests.With(prometheus.Labels{"url": url}).Inc()
}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	countEndpoint(request)
	io.WriteString(writer, "Hello world!\n")
}

func fooEndpoint(writer http.ResponseWriter, request *http.Request) {
	countEndpoint(request)
	io.WriteString(writer, "FOO!\n")
}

func barEndpoint(writer http.ResponseWriter, request *http.Request) {
	countEndpoint(request)
	io.WriteString(writer, "BAR!\n")
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
