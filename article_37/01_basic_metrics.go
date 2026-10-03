// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá sedmá část
//    Monitoring služeb a mikroslužeb psaných v Go nástrojem Prometheus
//    https://www.root.cz/clanky/monitoring-sluzeb-a-mikrosluzeb-psanych-v-go-nastrojem-prometheus/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Demonstrační příklad číslo 1:
//    HTTP server s endpointem s metrikami.
//
// Seznam demonstračních příkladů ze třicáté sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_37/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_37/01_basic_metrics.html
//

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	fmt.Println("Initializing HTTP server")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
