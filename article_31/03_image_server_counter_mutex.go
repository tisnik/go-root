// Seriál "Programovací jazyk Go"
//
// Třicátá první část
//
// Demonstrační příklad číslo 3:
//    Čítač měněný v mutexu.

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/rand"
	"net/http"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func indexPageHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	response := `
<body>
    <h1>Enterprise image renderer</h1>
    <img src="/image" width=256 height=256 />
</body>`
	io.WriteString(writer, response)
}

func calculateColor() color.RGBA {
	return color.RGBA{uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)), 255}
}

func imageHandler(writer http.ResponseWriter, request *http.Request) {
	const ImageWidth = 256
	const ImageHeight = 256
	outputimage := image.NewRGBA(image.Rectangle{image.Point{0, 0},
		image.Point{ImageWidth, ImageHeight}})

	for y := 0; y < ImageHeight; y++ {
		for x := 0; x < ImageWidth; x++ {
			c := calculateColor()
			outputimage.Set(x, y, c)
		}
	}
	png.Encode(writer, outputimage)
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func counterHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(writer, "Counter: %d\n", counter)
}

func main() {
	http.HandleFunc("/", indexPageHandler)
	http.HandleFunc("/image", imageHandler)
	http.HandleFunc("/counter", counterHandler)
	http.ListenAndServe(":8080", nil)
}
