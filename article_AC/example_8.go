package main

import (
	"fmt"
	"log"

	tcell "github.com/gdamore/tcell/v2"
)

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	column := x1
	for _, r := range []rune(text) {
		s.SetContent(column, row, r, nil, style)
		column++
		if column >= x2 {
			row++
			column = x1
		}
		if row > y2 {
			break
		}
	}
}

func drawPalette(screen tcell.Screen, green int32, offset int) {
	style := tcell.StyleDefault
	for j := 0; j < 16; j++ {
		for i := 0; i < 16; i++ {
			color := tcell.NewRGBColor(int32(i<<4), green, int32(j<<4))
			s := style.Foreground(tcell.Color(color))
			msg := fmt.Sprintf(" %02x%02x%02x ", i, green, j)
			drawText(screen, i*10, j+offset, i*10+8, j+offset, s, msg)
		}
	}
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	err = screen.Init()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	screen.SetStyle(tcell.StyleDefault)
	screen.Clear()

	defer func() {
		screen.Fini()
	}()

	drawPalette(screen, 0, 1)
	drawPalette(screen, 255, 20)

	for {
		screen.Show()

		event := screen.PollEvent()

		switch ev := event.(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				screen.Clear()
			}
		}
	}
}
