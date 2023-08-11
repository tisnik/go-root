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

func viewOnTop(gui *gocui.Gui, viewName string) error {
	gui.SetViewOnTop(viewName)
	return nil

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
		log.Fatal(err)
	}

	err = gui.SetKeybinding("", '1', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error { return viewOnTop(gui, "view1") })
	if err != nil {
		log.Fatal(err)
	}

	err = gui.SetKeybinding("", '2', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error { return viewOnTop(gui, "view2") })
	if err != nil {
		log.Fatal(err)
	}

	err = gui.SetKeybinding("", '3', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error { return viewOnTop(gui, "view3") })
	if err != nil {
		log.Fatal(err)
	}

	gui.SetCurrentView("")

	err = gui.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}
}
