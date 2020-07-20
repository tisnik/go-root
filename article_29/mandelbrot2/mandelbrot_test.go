// Seriál "Programovací jazyk Go"
//
// Dvacátá devátá část
//    Trasování a profilování aplikací naprogramovaných v Go
//    https://www.root.cz/clanky/trasovani-a-profilovani-aplikaci-naprogramovanych-v-go/

package renderer_test

import (
	"mandelbrot/renderer"
	"testing"
)

func TestRenderer(t *testing.T) {
	renderer.Start(256, 256, 255)
}
