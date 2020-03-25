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

	window, err := sdl.CreateWindow("Example #2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
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
	destRect.X = 160
	destRect.Y = 120
	destRect.W = 250
	destRect.H = 250

	err = convertedImage.BlitScaled(nil, primarySurface, &destRect)
	if err != nil {
		panic(err)
	}

	window.UpdateSurface()

	sdl.Delay(5000)
}
