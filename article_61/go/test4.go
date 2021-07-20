// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá první část
//    Programovací jazyk Go a 2D grafika
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-2d-grafika/

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

	window, err := sdl.CreateWindow("Example #4", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	primarySurface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	surfaceImage, err := img.Load("test1.bmp")
	if err != nil {
		panic(err)
	}
	defer surfaceImage.Free()

	convertedImage, err := surfaceImage.Convert(primarySurface.Format, 0)
	if err != nil {
		panic(err)
	}
	defer convertedImage.Free()

	primarySurface.FillRect(nil, sdl.MapRGB(convertedImage.Format, 192, 255, 192))

	var destRect sdl.Rect
	destRect.X = primarySurface.W/2 - convertedImage.W/4
	destRect.Y = primarySurface.H/2 - convertedImage.H/4
	destRect.W = convertedImage.W / 2
	destRect.H = convertedImage.H / 2

	err = convertedImage.BlitScaled(nil, primarySurface, &destRect)
	if err != nil {
		panic(err)
	}

	window.UpdateSurface()

	sdl.Delay(5000)
}
