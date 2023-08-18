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

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for column := x1; column <= x2; column++ {
			s.SetContent(column, row, ' ', nil, style)
		}
	}

	// Draw borders
	for column := x1; column <= x2; column++ {
		s.SetContent(column, y1, tcell.RuneHLine, nil, style)
		s.SetContent(column, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}

func drawBoxAroundScreen(screen tcell.Screen, style tcell.Style) {
	const offset = 5
	xmax, ymax := screen.Size()
	drawBox(screen, offset, offset, xmax-offset, ymax-offset, style, fmt.Sprintf("[%d, %d]", xmax, ymax))
}

func main() {
	defaultStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorDarkBlue)

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
	drawBoxAroundScreen(screen, boxStyle)

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
			// screen.Sync()
			screen.Clear()
			drawBoxAroundScreen(screen, boxStyle)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				screen.Clear()
			}
		}
	}
}
