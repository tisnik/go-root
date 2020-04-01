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
	Window           *sdl.Window
	PrimarySurface   *sdl.Surface
	Image            *sdl.Surface
	PlayerX, PlayerY int32
	NPCX, NPCY       int32
	NPCdX, NPCdY     int32
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

	state.Window, err = sdl.CreateWindow("SDL2 example #10", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
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

	state.PlayerX = 100
	state.PlayerY = 100
}

func (state *gameState) finalize() {
	state.PrimarySurface.Free()
	state.Image.Free()
	state.Window.Destroy()
	sdl.Quit()
}

func moveNPC(state *gameState) {
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

	playerRect := sdl.Rect{
		X: state.PlayerX,
		Y: state.PlayerY,
		W: 20,
		H: 20,
	}
	state.PrimarySurface.FillRect(&playerRect, sdl.MapRGB(state.PrimarySurface.Format, 255, 192, 192))

	moveNPC(state)

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
			state.PlayerX -= 2
		}
		if right {
			state.PlayerX += 2
		}
		if up {
			state.PlayerY -= 2
		}
		if down {
			state.PlayerY += 2
		}
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
