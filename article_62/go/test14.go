// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá druhá část
//    Programovací jazyk Go a 2D grafika – moduly sdl a img
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-2d-grafika-moduly-sdl-a-img/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z šedesáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_62/README.md

package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 640
	height = 480
)

func main() {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("SDL2 example #14", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	primarySurface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	primarySurface.FillRect(nil, sdl.MapRGB(primarySurface.Format, 192, 255, 192))

	dstRect := sdl.Rect{}
	dstRect.Y = 10

	primarySurface.Lock()
	var y int32
	for y = 0; y < primarySurface.H; y++ {
		scanLine := primarySurface.Pitch
		p := primarySurface.Pixels()
		offset := scanLine * y
		var x int32
		for x = 0; x < scanLine; x++ {
			p[offset+x] = byte(y)
		}
	}
	primarySurface.Unlock()

	window.UpdateSurface()

	sdl.Delay(5000)
}
