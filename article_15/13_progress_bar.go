package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
)

const BoardSize = 8

func CreateImage(width int, height int, imageIndex int) *image.Paletted {
	var palette = []color.Color{
		color.RGBA{50, 50, 50, 255},
		color.RGBA{0, 200, 200, 255},
	}

	if imageIndex == 0 {
		img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

		r := image.Rect(0, 0, width, height)
		draw.Draw(img, r, &image.Uniform{palette[0]}, image.ZP, draw.Src)

		return img
	} else {
		img := image.NewPaletted(image.Rect(imageIndex, 0, imageIndex+1, height), palette)

		r := image.Rect(imageIndex, 0, imageIndex+1, height)
		draw.Draw(img, r, &image.Uniform{palette[1]}, image.ZP, draw.Src)

		return img
	}

}

func main() {
	var images []*image.Paletted
	var delays []int
	var disposals []byte

	for i := 0; i < 200; i++ {
		img := CreateImage(200, 10, i)
		images = append(images, img)
		delays = append(delays, 2)
		disposals = append(disposals, gif.DisposalNone)
	}

	outfile, err := os.Create("13.gif")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	gif.EncodeAll(outfile, &gif.GIF{
		Image:    images,
		Delay:    delays,
		Disposal: disposals,
	})
}
