// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá první část
//    Tvorba grafů v jazyce Go: kreslení ve webovém klientu
//    https://www.root.cz/clanky/tvorba-grafu-v-jazyce-go-kresleni-ve-webovem-klientu/

package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Points struct {
	X []float64 `json:"x"`
	Y []float64 `json:"y"`
}

const npoints = 100
const periods = 2

func makePoints(npoints uint, offset float64) Points {
	var points Points
	points.X = make([]float64, npoints)
	points.Y = make([]float64, npoints)
	for i := uint(0); i < npoints; i++ {
		t := float64(i) * periods * 2.0 * math.Pi / float64(npoints)
		points.X[i] = t
		// limita
		if t == 0.0 {
			points.Y[i] = 1.0
		} else {
			points.Y[i] = math.Sin(t+offset) / t
		}
	}
	return points
}

func writePoints(writer http.ResponseWriter, points Points) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(points)
}

func seriesHandler(writer http.ResponseWriter, request *http.Request) {
	offsetStr := request.URL.Query().Get("offset")
	offset, err := strconv.ParseFloat(offsetStr, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	points := makePoints(npoints, offset)
	writePoints(writer, points)
}

func startHttpServer(address string) {
	log.Printf("Starting server on address %s", address)
	http.Handle("/", http.FileServer(http.Dir("plotly08/")))
	http.HandleFunc("/series", seriesHandler)
	http.ListenAndServe(address, nil)
}

func main() {
	startHttpServer(":8080")
}
