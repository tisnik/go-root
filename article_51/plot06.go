package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

const points = 50

const resX = 20.0 / 3.0 * vg.Inch
const resY = 5.0 * vg.Inch

func fillInSeries(offset float64) plotter.XYs {
	series := make(plotter.XYs, points)
	function1 := func(t float64, offset float64) float64 {
		// limita
		if t == 0.0 {
			return 1.0
		}
		return math.Sin(t-offset) / t
	}

	for i := range series {
		t := float64(i)*5.0*math.Pi/points + 0.4
		series[i].X = t
		series[i].Y = function1(t, offset)
	}
	return series
}

func main() {
	series1 := fillInSeries(0)
	series2 := fillInSeries(math.Pi / 2.0)
	series3 := fillInSeries(-math.Pi / 2.0)

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Points Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(series1)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

	l, err := plotter.NewLine(series2)
	if err != nil {
		panic(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	lpLine, lpPoints, err := plotter.NewLinePoints(series3)
	if err != nil {
		panic(err)
	}
	lpLine.Color = color.RGBA{G: 255, A: 255}
	lpPoints.Shape = draw.PyramidGlyph{}
	lpPoints.Color = color.RGBA{R: 255, A: 255}

	p.Add(s, l, lpLine, lpPoints)
	p.Legend.Add("scatter", s)
	p.Legend.Add("line", l)
	p.Legend.Add("line points", lpLine, lpPoints)

	err = p.Save(resX, resY, "plot06.png")
	if err != nil {
		panic(err)
	}
}
