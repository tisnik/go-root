package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 640
	height = 480
)

type gameState struct {
	Window         *sdl.Window
	PrimarySurface *sdl.Surface
	Image          *sdl.Surface
}

func initialize(state *gameState) {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	state.Window, err = sdl.CreateWindow("SDL2 example #4", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	state.PrimarySurface, err = state.Window.GetSurface()
	if err != nil {
		panic(err)
	}

	state.Image, err = img.Load("globe.png")
	if err != nil {
		panic(err)
	}
}

func finalize(state *gameState) {
	state.PrimarySurface.Free()
	state.Image.Free()
	state.Window.Destroy()
	sdl.Quit()
}

func redraw(state *gameState) {
	state.PrimarySurface.FillRect(nil, sdl.MapRGB(state.PrimarySurface.Format, 192, 255, 192))

	dstRect := sdl.Rect{
		X: width/2 - state.Image.W/2,
		Y: height/2 - state.Image.H/2,
		W: 0,
		H: 0,
	}

	err := state.Image.Blit(nil, state.PrimarySurface, &dstRect)
	if err != nil {
		panic(err)
	}

	state.Window.UpdateSurface()
}

func main() {
	var state gameState
	initialize(&state)
	defer finalize(&state)

	sdl.Delay(10)
	redraw(&state)
	sdl.Delay(5000)
}
