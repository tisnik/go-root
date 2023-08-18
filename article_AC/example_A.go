package main

import (
	"log"

	tcell "github.com/gdamore/tcell/v2"
)

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
	screen.ShowCursor(10, 5)

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
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			}
			switch ev.Rune() {
			case 'C':
				fallthrough
			case 'c':
				screen.Clear()
			case '1':
				screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
			case '2':
				screen.SetCursorStyle(tcell.CursorStyleSteadyBlock)
			case '3':
				screen.SetCursorStyle(tcell.CursorStyleBlinkingUnderline)
			case '4':
				screen.SetCursorStyle(tcell.CursorStyleSteadyUnderline)
			case '5':
				screen.SetCursorStyle(tcell.CursorStyleBlinkingBar)
			case '6':
				screen.SetCursorStyle(tcell.CursorStyleSteadyBar)
			}
		}
	}
}
