package main

import (
	_ "embed"

	. "modernc.org/tk9.0"
	_ "modernc.org/tk9.0/themes/azure"
)

//go:embed gopher.png
var gopher []byte

func main() {
	ActivateTheme("azure light")
	Pack(Label(Image(NewPhoto(Data(gopher)))),
		TExit(),
		Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
	App.Center().Wait()
}
