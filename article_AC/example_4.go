package main

import (
	"log"

	tcell "github.com/gdamore/tcell/v2"
)

func drawStar(s tcell.Screen, x, y int, style tcell.Style) {
	s.SetContent(x, y, '*', nil, style)
}

func main() {
	defaultStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorReset)
	starStyle := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	err = screen.Init()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	screen.EnableMouse()
	screen.SetStyle(defaultStyle)
	screen.Clear()

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
		case *tcell.EventMouse:
			if ev.Buttons() == tcell.Button1 {
				x, y := ev.Position()
				drawStar(screen, x, y, starStyle)
			}

		}
	}
}
