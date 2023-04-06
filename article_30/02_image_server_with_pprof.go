// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá část
//    Trasování a profilování aplikací naprogramovaných v Go (dokončení
//    https://www.root.cz/clanky/trasovani-a-profilovani-aplikaci-naprogramovanych-v-go-dokonceni/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze třicáté části:
//    https://github.com/tisnik/go-root/blob/master/article_30/README.md
//
// Demonstrační příklad číslo 2:
//     HTTP server nabízející i pprof metriky.

package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
)

func indexPageHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	response := `
<body>
    <h1>Enterprise image renderer</h1>
    <img src="/image" width=256 height=256 />
</body>`
	writer.Write([]byte(response))
}

func calculateColor() color.RGBA {
	return color.RGBA{uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)), 255}
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
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
	png.Encode(w, outputimage)
}

func main() {
	http.HandleFunc("/", indexPageHandler)
	http.HandleFunc("/image", imageHandler)
	http.ListenAndServe(":8080", nil)
}
