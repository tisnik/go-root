package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func layoutManager(gui *gocui.Gui) error {
	view1, err := gui.SetView("view1", 10, 5, 20, 10)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(view1, "view1")
	}

	view2, err := gui.SetView("view2", 14, 7, 24, 12)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(view2, "view2")
	}

	view3, err := gui.SetView("view3", 18, 9, 28, 14)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(view3, "view3")
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

	err = gui.SetKeybinding("", 'q', gocui.ModNone, quitEvent)
	if err != nil {
		log.Print(err)
		return
	}

	gui.SetCurrentView("")

	err = gui.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Print(err)
		return
	}
}
