package main

import (
	"machine"
)

var pins [8]machine.Pin

func displaySegments(bits []bool) {
	for i := range bits {
		bit := bits[i]
		pin := pins[i]
		if bit {
			pin.High()
		} else {
			pin.Low()
		}
	}
}

func main() {
	c1 := machine.GP5
	c1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	c1.High()

	pins[0] = machine.GP11
	pins[1] = machine.GP9
	pins[2] = machine.GP13
	pins[3] = machine.GP15
	pins[4] = machine.GP16
	pins[5] = machine.GP10
	pins[6] = machine.GP12
	pins[7] = machine.GP14

	for _, pin := range pins {
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	displaySegments([]bool{true, true, true, true, true, true, true, true})
}
