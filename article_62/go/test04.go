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
	"github.com/veandco/go-sdl2/img"
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

	window, err := sdl.CreateWindow("SDL2 example #4", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	primarySurface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	tempImage, err := img.Load("test1.bmp")
	if err != nil {
		panic(err)
	}
	defer tempImage.Free()

	convertedImage, err := tempImage.Convert(primarySurface.Format, 0)
	if err != nil {
		panic(err)
	}
	defer convertedImage.Free()

	primarySurface.FillRect(nil, sdl.MapRGB(primarySurface.Format, 192, 255, 192))

	err = convertedImage.Blit(nil, primarySurface, nil)
	if err != nil {
		panic(err)
	}

	window.UpdateSurface()

	sdl.Delay(5000)
}
