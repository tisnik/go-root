// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z šedesáté čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_64/README.md

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
	NPCX, NPCY     int32
	NPCdX, NPCdY   int32
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

	state.Window, err = sdl.CreateWindow("Game #1", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
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

	state.NPCX = width/2 - state.Image.W/2
	state.NPCY = height/2 - state.Image.H/2
	state.NPCdX = 1
	state.NPCdY = 1
}

func (state *gameState) finalize() {
	if state.PrimarySurface != nil {
		state.PrimarySurface.Free()
	}

	if state.Image != nil {
		state.Image.Free()
	}

	if state.Window != nil {
		state.Window.Destroy()
	}
	sdl.Quit()
}

func (state *gameState) moveNPC() {
	const (
		LeftBorder   = 5
		TopBorder    = 5
		RightBorder  = 10
		BottomBorder = 10
	)
	state.NPCX += state.NPCdX
	state.NPCY += state.NPCdY

	if state.NPCX > state.PrimarySurface.W-state.Image.W+RightBorder {
		state.NPCdX = -state.NPCdX
	}
	if state.NPCY > state.PrimarySurface.H-state.Image.H+BottomBorder {
		state.NPCdY = -state.NPCdY
	}
	if state.NPCX < -LeftBorder {
		state.NPCdX = -state.NPCdX
	}
	if state.NPCY < -TopBorder {
		state.NPCdY = -state.NPCdY
	}
}

func (state *gameState) redraw() {
	state.PrimarySurface.FillRect(nil, sdl.MapRGB(state.PrimarySurface.Format, 192, 255, 192))

	dstRect := sdl.Rect{
		X: state.NPCX,
		Y: state.NPCY,
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
				}
			}
		}
		state.moveNPC()
		state.redraw()
		sdl.Delay(10)
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
