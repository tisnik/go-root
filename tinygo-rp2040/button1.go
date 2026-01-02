package main

import (
	"machine"
)

const InputPin = 0
const OutputPin = 2

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
