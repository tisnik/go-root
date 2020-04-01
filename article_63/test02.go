package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 640
	height = 480
)

var (
	window         *sdl.Window
	primarySurface *sdl.Surface
	image          *sdl.Surface
)

func initialize() {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	window, err = sdl.CreateWindow("SDL2 example #2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	primarySurface, err = window.GetSurface()
	if err != nil {
		panic(err)
	}

	image, err = img.Load("globe.png")
	if err != nil {
		panic(err)
	}
}

func finalize() {
	primarySurface.Free()
	image.Free()
	window.Destroy()
	sdl.Quit()
}

func redraw() {
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
	initialize()
	defer finalize()

	sdl.Delay(10)
	redraw()
	sdl.Delay(5000)
}
