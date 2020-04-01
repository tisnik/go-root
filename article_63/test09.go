package main

import (
	"log"

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
	eventLoop()
}

type gameState struct {
	Window         *sdl.Window
	PrimarySurface *sdl.Surface
	Image          *sdl.Surface
	X, Y           int32
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

	state.Window, err = sdl.CreateWindow("SDL2 example #9", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
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

	state.X = width/2 - state.Image.W/2
	state.Y = height/2 - state.Image.H/2
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
		X: state.X,
		Y: state.Y,
		W: 0,
		H: 0,
	}

	err := state.Image.Blit(nil, state.PrimarySurface, &dstRect)
	if err != nil {
		panic(err)
	}

	state.Window.UpdateSurface()
}

func (state *gameState) eventLoop() {
	var event sdl.Event
	done := false
	left := false
	right := false
	up := false
	down := false

	for !done {
		event = sdl.PollEvent()
		switch t := event.(type) {
		case *sdl.QuitEvent:
			done = true
		case *sdl.KeyboardEvent:
			keyCode := t.Keysym.Sym
			switch t.State {
			case sdl.PRESSED:
				switch keyCode {
				case sdl.K_ESCAPE:
					done = true
				case sdl.K_q:
					done = true
				case sdl.K_LEFT:
					left = true
				case sdl.K_RIGHT:
					right = true
				case sdl.K_UP:
					up = true
				case sdl.K_DOWN:
					down = true
				}
			case sdl.RELEASED:
				switch keyCode {
				case sdl.K_LEFT:
					left = false
				case sdl.K_RIGHT:
					right = false
				case sdl.K_UP:
					up = false
				case sdl.K_DOWN:
					down = false
				}
			}
		}
		if left {
			state.X -= 2
		}
		if right {
			state.X += 2
		}
		if up {
			state.Y -= 2
		}
		if down {
			state.Y += 2
		}
		state.redraw()
		sdl.Delay(20)
	}
	log.Println("Quitting")
}

func main() {
	game := NewGame()
	game.initialize()
	defer game.finalize()

	sdl.Delay(10)
	game.eventLoop()
}
