package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 320
	height = 240
)

var pixmap *sdl.Surface = nil
var primarySurface *sdl.Surface = nil
var window *sdl.Window = nil

var xpos float64 = -0.75
var ypos float64 = 0.0
var scale float64 = 150.0
var uhel float64 = 45.0

func min(a, b int32) int32 {
	if a > b {
		return b
	}
	return a
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func gfxInitialize(width int32, height int32, bpp int) {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	window, err = sdl.CreateWindow("Fractals", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	primarySurface, err = window.GetSurface()
	if err != nil {
		panic(err)
	}
}

func gfxFinalize() {
	primarySurface.Free()
	window.Destroy()
	sdl.Quit()
}

func createSurface(width int32, height int32) *sdl.Surface {
	surface, err := sdl.CreateRGBSurface(sdl.SWSURFACE, width, height, 32, 0x00ff0000, 0x0000ff00, 0x000000ff, 0x00000000)
	if err != nil {
		panic(err)
	}
	return surface
}

func putpixel(surface *sdl.Surface, x int32, y int32, r byte, g byte, b byte) {
	if x >= 0 && x < surface.W && y >= 0 && y < surface.H {
		switch surface.Format.BitsPerPixel {
		case 24:
			index := x*3 + y*surface.Pitch
			pixels := surface.Pixels()
			pixels[index] = b
			pixels[index+1] = g
			pixels[index+2] = r
		case 32:
			index := x*4 + y*surface.Pitch
			pixels := surface.Pixels()
			pixels[index] = b
			pixels[index+1] = g
			pixels[index+2] = r
		}
	}
}

func hline(surface *sdl.Surface, x1 int32, x2 int32, y int32, r byte, g byte, b byte) {
	fromX := min(x1, x2)
	toX := max(x1, x2)

	for x := fromX; x <= toX; x++ {
		putpixel(surface, x, y, r, g, b)
	}
}

func vline(surface *sdl.Surface, x int32, y1 int32, y2 int32, r byte, g byte, b byte) {
	fromY := min(y1, y2)
	toY := max(y1, y2)

	for y := fromY; y <= toY; y++ {
		putpixel(surface, x, y, r, g, b)
	}
}

func drawGrid(surface *sdl.Surface) {
	width := surface.W
	height := surface.H
	var x, y int32

	surface.FillRect(nil, 0xffffffff)

	for x = 0; x < width; x += 20 {
		vline(surface, x, 0, height-1, 191, 191, 255)
	}

	for y = 0; y < height; y += 20 {
		hline(surface, 0, width-1, y, 191, 191, 255)
	}
}

func showFractal(surface *sdl.Surface) {
	err := surface.Blit(nil, primarySurface, nil)
	if err != nil {
		panic(err)
	}
	window.UpdateSurface()
}

func calcCorner(xpos, ypos, scale float64) (xmin, ymin, xmax, ymax float64) {
	xmin = xpos - width/scale
	ymin = ypos - height/scale
	xmax = xpos + width/scale
	ymax = ypos + height/scale
	return
}

func drawMandelJulia(surface *sdl.Surface) {
	ccx := 0.0
	ccy := 0.0

	xmin, ymin, xmax, ymax := calcCorner(xpos, ypos, scale)

	u := uhel * 3.1415 / 180.0
	cosu := math.Cos(u)
	sinu := math.Sin(u)
	ccxc := ccx * cosu
	ccyc := ccy * cosu

	cy := ymin
	for y := 0; y < height; y++ {
		cx := xmin
		for x := 0; x < width; x++ {
			i := 0
			zx := cx * cosu
			zy := cy * cosu
			for {
				zx2 := zx * zx
				zy2 := zy * zy
				zy = 2.0*zx*zy + ccyc + cy*sinu
				zx = zx2 - zy2 + ccxc + cx*sinu
				i++
				if i >= 64 || zx2+zy2 >= 4.0 {
					break
				}
			}
			b := i * 2
			g := i * 3
			r := i * 5
			putpixel(surface, int32(x+width/2), int32(y+height/2), byte(r), byte(g), byte(b))
			cx += (xmax - xmin) / width
		}
		cy += (ymax - ymin) / height
	}
}

func redraw(pixmap *sdl.Surface) {
	drawGrid(pixmap)
	drawMandelJulia(pixmap)
	showFractal(pixmap)
}

func mainEventLoop() {
	var event sdl.Event
	done := false
	left := false
	right := false
	up := false
	down := false
	zoomin := false
	zoomout := false
	angle1 := false
	angle2 := false

	for !done {
		event = sdl.PollEvent()
		switch t := event.(type) {
		case *sdl.QuitEvent:
			done = true
		case *sdl.KeyboardEvent:
			keyCode := t.Keysym.Sym
			if t.State == sdl.PRESSED {
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
				case sdl.K_PAGEDOWN:
					zoomin = true
				case sdl.K_PAGEUP:
					zoomout = true
				case sdl.K_z:
					angle1 = true
				case sdl.K_x:
					angle2 = true
				}
			}
			if t.State == sdl.RELEASED {
				switch keyCode {
				case sdl.K_LEFT:
					left = false
				case sdl.K_RIGHT:
					right = false
				case sdl.K_UP:
					up = false
				case sdl.K_DOWN:
					down = false
				case sdl.K_PAGEDOWN:
					zoomin = false
				case sdl.K_PAGEUP:
					zoomout = false
				case sdl.K_z:
					angle1 = false
				case sdl.K_x:
					angle2 = false
				}
			}
		}
		performRedraw := false
		if left {
			xpos -= 10.0 / scale
			performRedraw = true
		}
		if right {
			xpos += 10.0 / scale
			performRedraw = true
		}
		if up {
			ypos -= 10.0 / scale
			performRedraw = true
		}
		if down {
			ypos += 10.0 / scale
			performRedraw = true
		}
		if zoomin {
			scale *= 0.9
			performRedraw = true
		}
		if zoomout {
			scale *= 1.1
			performRedraw = true
		}
		if angle1 {
			uhel--
			performRedraw = true
		}
		if angle2 {
			uhel++
			performRedraw = true
		}
		if performRedraw {
			redraw(pixmap)
			sdl.Delay(20)
		}
	}
}

func main() {
	gfxInitialize(640, 480, 32)
	defer gfxFinalize()

	pixmap = createSurface(primarySurface.W, primarySurface.H)
	redraw(pixmap)
	mainEventLoop()
}
