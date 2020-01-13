package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
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

func series1Handler(writer http.ResponseWriter, request *http.Request) {
	points := makePoints(npoints, 0)
	writePoints(writer, points)
}

func series2Handler(writer http.ResponseWriter, request *http.Request) {
	points := makePoints(npoints, npoints/2.0)
	writePoints(writer, points)
}

func series3Handler(writer http.ResponseWriter, request *http.Request) {
	points := makePoints(npoints, -npoints/2.0)
	writePoints(writer, points)
}

func startHttpServer(address string) {
	log.Printf("Starting server on address %s", address)
	http.Handle("/", http.FileServer(http.Dir("plotly06/")))
	http.HandleFunc("/series1", series1Handler)
	http.HandleFunc("/series2", series2Handler)
	http.HandleFunc("/series3", series3Handler)
	http.ListenAndServe(address, nil)
}

func main() {
	startHttpServer(":8080")
}
