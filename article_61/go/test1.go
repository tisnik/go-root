// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá první část
//    Programovací jazyk Go a 2D grafika
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-2d-grafika/
//
// Seznam příkladů z šedesáté první části:
//    https://github.com/tisnik/go-root/blob/master/article_61/README.md

package main

import "github.com/veandco/go-sdl2/sdl"

const (
	width  = 640
	height = 480
)

func main() {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Example #1", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
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
	window.UpdateSurface()

	sdl.Delay(5000)
}
