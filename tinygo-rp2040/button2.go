package main

import (
	"machine"
)

const InputPin = machine.GP0
const OutputPin = machine.GP2

func main() {
	led := machine.Pin(OutputPin)
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button := machine.Pin(InputPin)
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		pressed := button.Get()
		if pressed {
			led.High()
		} else {
			led.Low()
		}
	}
}
