// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá třetí část
//    Programovací jazyk Go a 2D grafika – kostra jednoduché hry
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-2d-grafika-kostra-jednoduche-hry/

package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 640
	height = 480
)

func initialize() (*sdl.Window, *sdl.Surface, *sdl.Surface) {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("SDL2 example #3", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	primarySurface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	image, err := img.Load("globe.png")
	if err != nil {
		panic(err)
	}

	return window, primarySurface, image
}

func finalize(window *sdl.Window, primarySurface *sdl.Surface, image *sdl.Surface) {
	primarySurface.Free()
	image.Free()
	window.Destroy()
	sdl.Quit()
}

func redraw(window *sdl.Window, primarySurface *sdl.Surface, image *sdl.Surface) {
	primarySurface.FillRect(nil, sdl.MapRGB(primarySurface.Format, 192, 255, 192))

	dstRect := sdl.Rect{
		X: width/2 - image.W/2,
		Y: height/2 - image.H/2,
		W: 0,
		H: 0,
	}

	err := image.Blit(nil, primarySurface, &dstRect)
	if err != nil {
		panic(err)
	}

	window.UpdateSurface()
}

func main() {
	window, primarySurface, image := initialize()
	defer finalize(window, primarySurface, image)

	sdl.Delay(10)
	redraw(window, primarySurface, image)
	sdl.Delay(5000)
}
