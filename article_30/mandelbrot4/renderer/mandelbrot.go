// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá část
//    Trasování a profilování aplikací naprogramovaných v Go (dokončení
//    https://www.root.cz/clanky/trasovani-a-profilovani-aplikaci-naprogramovanych-v-go-dokonceni/

package renderer

import (
	"context"
	"image"
	"image/png"
	"log"
	"os"
	"runtime/trace"
)

func writeImage(width uint, height uint, pixels []byte) {
	img := image.NewNRGBA(image.Rect(0, 0, int(width), int(height)))
	pixel := 0

	for y := 0; y < int(height); y++ {
		offset := img.PixOffset(0, y)
		for x := uint(0); x < width; x++ {
			img.Pix[offset] = pixels[pixel]
			img.Pix[offset+1] = pixels[pixel+1]
			img.Pix[offset+2] = pixels[pixel+2]
			img.Pix[offset+3] = 0xff
			pixel += 3
			offset += 4
		}
	}

	outputFile, err := os.Create("mandelbrot.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	png.Encode(outputFile, img)
}

func iterCount(cx float64, cy float64, maxiter uint) uint {
	var zx float64 = 0.0
	var zy float64 = 0.0
	var i uint = 0
	for i < maxiter {
		zx2 := zx * zx
		zy2 := zy * zy
		if zx2+zy2 > 4.0 {
			break
		}
		zy = 2.0*zx*zy + cy
		zx = zx2 - zy2 + cx
		i++
	}
	return i
}

func calcMandelbrot(width uint, height uint, maxiter uint, palette [][3]byte, image []byte, cy float64, done chan bool) {
	var cx float64 = -2.0
	for x := uint(0); x < width; x++ {
		i := iterCount(cx, cy, maxiter)
		color := palette[i]
		image[3*x] = color[0]
		image[3*x+1] = color[1]
		image[3*x+2] = color[2]
		cx += 3.0 / float64(width)
	}
	done <- true
}

func Start(ctx context.Context, width int, height int, maxiter int) {
	trace.Logf(ctx, "settings", "width %i", width)
	trace.Logf(ctx, "settings", "height %i", height)
	trace.Logf(ctx, "settings", "maxiter %i", maxiter)

	done := make(chan bool, height)

	pixels := make([]byte, width*height*3)
	offset := 0
	delta := width * 3

	var cy float64 = -1.5

	region := trace.StartRegion(ctx, "startGoroutines")
	for y := 0; y < height; y++ {
		trace.Logf(ctx, "y (scanline)", "%i", y)
		go calcMandelbrot(uint(width), uint(height), uint(maxiter), mandmap[:], pixels[offset:offset+delta], cy, done)
		offset += delta
		cy += 3.0 / float64(height)
	}
	region.End()

	region2 := trace.StartRegion(ctx, "waitForGoroutines")
	for i := 0; i < height; i++ {
		<-done
	}
	region2.End()

	region3 := trace.StartRegion(ctx, "writeImage")
	writeImage(uint(width), uint(height), pixels)
	region3.End()
}
