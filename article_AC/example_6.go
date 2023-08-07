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

	style := tcell.StyleDefault
	for i := 0; i < 16; i++ {
		s := style.Foreground(tcell.PaletteColor(i))
		msg := fmt.Sprintf("Color #%d", i)
		drawText(screen, 5, i, 30, i, s, msg)
	}

	defer func() {
		screen.Fini()
	}()

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
