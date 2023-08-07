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
	defaultStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	err = screen.Init()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	screen.SetStyle(defaultStyle)
	screen.Clear()
	drawText(screen, 5, 5, 20, 20, defaultStyle, "Hello, world!")

	quit := func() {
		maybePanic := recover()
		screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

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
