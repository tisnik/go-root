// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá čtvrtá část
//    Programovací jazyk Go a 2D grafika – kostra jednoduché hry
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-2d-grafika-kostra-jednoduche-hry/
//
// Seznam demonstračních příkladů z šedesáté čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_64/README.md

package main

import (
	"fmt"
	"log"
	"time"

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

type NPCState struct {
	Texture *sdl.Texture
	Width   int32
	Height  int32
	X, Y    int32
	DX, DY  int32
}

type gameState struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
	NPC      NPCState
	frames   int
}

func NewGame() gameState {
	return gameState{
		Window: nil,
		frames: 0,
	}
}

func (state *gameState) initialize() {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	state.Window, err = sdl.CreateWindow("Game #3", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	state.Renderer, err = sdl.CreateRenderer(state.Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	image, err := img.Load("globe.png")
	if err != nil {
		panic(err)
	}
	defer image.Free()

	state.NPC.Texture, err = state.Renderer.CreateTextureFromSurface(image)
	if err != nil {
		panic(err)
	}

	state.NPC.X = width/2 - image.W/2
	state.NPC.Y = height/2 - image.H/2
	state.NPC.Width = image.W
	state.NPC.Height = image.H
	state.NPC.DX = 1
	state.NPC.DY = 1
}

func (state *gameState) finalize() {
	if state.NPC.Texture != nil {
		state.NPC.Texture.Destroy()
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
	state.NPC.X += state.NPC.DX
	state.NPC.Y += state.NPC.DY

	width, height := state.Window.GetSize()

	if state.NPC.X > width-state.NPC.Width+RightBorder {
		state.NPC.DX = -state.NPC.DX
	}
	if state.NPC.Y > height-state.NPC.Height+BottomBorder {
		state.NPC.DY = -state.NPC.DY
	}
	if state.NPC.X < -LeftBorder {
		state.NPC.DX = -state.NPC.DX
	}
	if state.NPC.Y < -TopBorder {
		state.NPC.DY = -state.NPC.DY
	}
}

func (state *gameState) redraw() {
	state.frames++
	state.Renderer.SetDrawColor(192, 255, 192, 255)
	state.Renderer.Clear()

	dstRect := sdl.Rect{
		X: state.NPC.X,
		Y: state.NPC.Y,
		W: state.NPC.Width,
		H: state.NPC.Height,
	}

	err := state.Renderer.Copy(state.NPC.Texture, nil, &dstRect)
	if err != nil {
		panic(err)
	}

	state.Renderer.Present()
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
	startTime := time.Now()

	game.eventLoop()

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)

	framerate := int(float64(game.frames) / elapsed.Seconds())
	fmt.Printf("Frames: %d Framerate %d\n", game.frames, framerate)
}
