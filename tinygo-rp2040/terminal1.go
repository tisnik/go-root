package main

import (
	"machine"
	"time"
)

const InputPin = machine.GP0
const OutputPin = machine.GP2

func main() {
	led := OutputPin
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button := InputPin
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		pressed := button.Get()
		led.Set(pressed)
		if pressed {
			println("Button press detected")
			const SleepAmount = time.Second * 1
			time.Sleep(SleepAmount)
		}
	}
}
