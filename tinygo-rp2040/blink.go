package main

import (
	"machine"
	"time"
)

const OutputPin = 2
const SleepAmount = time.Millisecond * 200

func main() {
	led := machine.Pin(OutputPin)
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.High()
		time.Sleep(SleepAmount)
		led.Low()
		time.Sleep(SleepAmount)
	}
}
