package main

import (
	"machine"
	"time"
)

const SleepAmount = time.Millisecond * 1

var controls [4]machine.Pin
var pins [8]machine.Pin

var digits [][]bool = [][]bool{
	{true, true, true, true, true, true, false, false},
	{false, true, true, false, false, false, false, false},
	{true, true, false, true, true, false, true, false},
	{true, true, true, true, false, false, true, false},
	{false, true, true, false, false, true, true, false},
	{true, false, true, true, false, true, true, false},
	{true, false, true, true, true, true, true, false},
	{true, true, true, false, false, false, false, false},
	{true, true, true, true, true, true, true, false},
	{true, true, true, true, false, true, true, false},
}

func init() {
	controls[0] = machine.GP5
	controls[1] = machine.GP6
	controls[2] = machine.GP7
	controls[3] = machine.GP8
	for _, control := range controls {
		control.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

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
}

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

func displayNumber(number int) {
	x := number
	for i := range 4 {
		digit := x % 10
		x /= 10
		bits := digits[digit]
		control := controls[3-i]
		control.High()
		displaySegments(bits)
		time.Sleep(SleepAmount)
		control.Low()
	}
}

func main() {
	for {
		temperature := machine.ReadTemperature()
		temperature /= 10
		displayNumber(int(temperature))
	}
}
