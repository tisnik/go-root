package renderer_test

import (
	"mandelbrot/renderer"
	"testing"
)

func TestRenderer(t *testing.T) {
	renderer.Start(256, 256, 255)
}
