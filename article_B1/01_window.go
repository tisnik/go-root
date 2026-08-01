package main

import (
	"fmt"

	"github.com/gogpu/gg"
	_ "github.com/gogpu/gg/gpu"
	"github.com/gogpu/gogpu"
)

func main() {
	// rozměry okna
	const width = 400
	const height = 300

	// inicializace aplikace s GUI
	app := gogpu.NewApp(gogpu.DefaultConfig().
		WithTitle("GoGPU + gg").
		WithSize(width, height))

	// callback metoda zavolaná ve chvíli, kdy se má překreslit okno
	app.OnDraw(func(dc *gogpu.Context) {
		fmt.Println("Redraw....")
	})

	// callback metoda zavolaná ve chvíli, kdy se okno uzavírá
	app.OnClose(func() {
		gg.CloseAccelerator()
	})

	// spuštění aplikace
	app.Run()
}
