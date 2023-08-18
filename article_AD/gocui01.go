package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func layoutManager(gui *gocui.Gui) error {
	const message = "www.root.cz"

	width := len(message) + 1
	view, err := gui.SetView("mainView", 10, 5, 10+width, 7)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(view, message)
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

	err = gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quitEvent)
	if err != nil {
		log.Print(err)
		return
	}

	err = gui.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Print(err)
		return
	}
}
