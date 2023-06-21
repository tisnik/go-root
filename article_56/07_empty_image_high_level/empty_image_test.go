// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá šestá část
//    Programovací jazyk Go a assembler (3.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-assembler-3-cast/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z padesáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_56/README.md

package main

import (
	"fmt"
	"image"
	"testing"
)

var sizes = []int{32, 128, 256, 512, 1024, 2048}

func BenchmarkFillPixels(b *testing.B) {
	for _, size := range sizes {
		sizeStr := fmt.Sprintf("%dx%d", size, size)
		b.Run(sizeStr, func(b *testing.B) {
			destinationImage := image.NewRGBA(image.Rect(0, 0, size, size))

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				fillPixels(destinationImage)
			}
		})
	}
}
