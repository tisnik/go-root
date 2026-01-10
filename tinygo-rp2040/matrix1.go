package main

import (
	"machine"
	"time"
)

func main() {
	rowsLatch := machine.GP19
	rowsData := machine.GP18
	rowsClock := machine.GP17

	colsLatch := machine.GP22
	colsData := machine.GP21
	colsClock := machine.GP20

	rowsLatch.Configure(machine.PinConfig{Mode: machine.PinOutput})
	rowsData.Configure(machine.PinConfig{Mode: machine.PinOutput})
	rowsClock.Configure(machine.PinConfig{Mode: machine.PinOutput})

	colsLatch.Configure(machine.PinConfig{Mode: machine.PinOutput})
	colsData.Configure(machine.PinConfig{Mode: machine.PinOutput})
	colsClock.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		colsData.Low()
		colsClock.High()
		colsData.High()
		for x := range 32 {
			for y := range 8 {
				rowsData.Set(true)
				rowsClock.Low()
				time.Sleep(1 * time.Millisecond)
				rowsClock.High()
			}
			colsClock.High()
			time.Sleep(1 * time.Millisecond)
			colsClock.Low()
			colsLatch.High()
			time.Sleep(1 * time.Millisecond)
			colsLatch.Low()
			rowsLatch.High()
			time.Sleep(1 * time.Millisecond)
			rowsLatch.Low()
		}
		time.Sleep(10 * time.Millisecond)
	}
}
