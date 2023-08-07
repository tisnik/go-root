package main

import (
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
	drawText(screen, 5, 1, 30, 1, style, "Normal text")
	drawText(screen, 5, 2, 30, 2, style.Bold(true), "Bold text")
	drawText(screen, 5, 3, 30, 3, style.Italic(true), "Italic text")
	drawText(screen, 5, 4, 30, 4, style.Underline(true), "Underline text")
	drawText(screen, 5, 5, 30, 5, style.StrikeThrough(true), "Strike through text")
	drawText(screen, 5, 6, 30, 6, style.Blink(true), "Blink")
	drawText(screen, 5, 7, 30, 7, style.Reverse(true), "Reverse")
	drawText(screen, 5, 8, 30, 8, style.Url("https://www.root.cz"), "https://www.root.cz")

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
