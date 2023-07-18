// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá třetí část
//    Programovací jazyk Go a 2D grafika – kostra jednoduché hry
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-2d-grafika-kostra-jednoduche-hry/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z šedesáté třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_63/README.md

package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 640
	height = 480
)

type Game interface {
	initialize()
	finalize()
	redraw()
}

type gameState struct {
	Window         *sdl.Window
	PrimarySurface *sdl.Surface
	Image          *sdl.Surface
}

func NewGame() gameState {
	return gameState{
		Window:         nil,
		PrimarySurface: nil,
		Image:          nil,
	}
}

func (state *gameState) initialize() {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	state.Window, err = sdl.CreateWindow("SDL2 example #5", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
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

func (state *gameState) finalize() {
	state.PrimarySurface.Free()
	state.Image.Free()
	state.Window.Destroy()
	sdl.Quit()
}

func (state *gameState) redraw() {
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
	game := NewGame()
	game.initialize()
	defer game.finalize()

	sdl.Delay(10)
	game.redraw()
	sdl.Delay(5000)
}
