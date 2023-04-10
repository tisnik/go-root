// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá první část
//    Sledování vybraných metrik služeb naprogramovaných v jazyku Go
//    https://www.root.cz/clanky/sledovani-vybranych-metrik-sluzeb-naprogramovanych-v-jazyku-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze třicáté první části:
//    https://github.com/tisnik/go-root/blob/master/article_31/README.md
//
// Demonstrační příklad číslo 6:
//    Úprava předchozího příkladu používajícího uzávěry.

package main

import (
	"expvar"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/rand"
	"net/http"
)

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

func imageHandler(writer http.ResponseWriter, request *http.Request, counter *expvar.Int) {
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
	counter.Add(1)
}

func counterHandler(writer http.ResponseWriter, request *http.Request, counter *expvar.Int) {
	writer.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(writer, "Counter: %d\n", counter.Value())
}

func main() {
	counter := expvar.NewInt("counter")

	http.HandleFunc("/", indexPageHandler)
	http.HandleFunc("/image",
		func(writer http.ResponseWriter, request *http.Request) { imageHandler(writer, request, counter) })
	http.HandleFunc("/counter",
		func(writer http.ResponseWriter, request *http.Request) { counterHandler(writer, request, counter) })
	http.ListenAndServe(":8080", nil)
}
