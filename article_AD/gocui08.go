package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func layoutManager(gui *gocui.Gui) error {
	const message = "www.root.cz"

	maxX, maxY := gui.Size()

	width := len(message) + 1
	view, err := gui.SetView("mainView", maxX/2-width/2, maxY/2-1, maxX/2+width/2, maxY/2+1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(view, "\x1b[0;31m"+message)
	}
	return nil
}

func quitEvent(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	gui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal(err)
	}
	defer gui.Close()

	gui.SetManagerFunc(layoutManager)
	gui.SetCurrentView("")

	err = gui.SetKeybinding("", 'q', gocui.ModNone, quitEvent)
	if err != nil {
		log.Fatal(err)
	}

	err = gui.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}
}
