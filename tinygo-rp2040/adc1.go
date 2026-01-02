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
	controls[0] = machine.Pin(machine.GP5)
	controls[1] = machine.Pin(machine.GP6)
	controls[2] = machine.Pin(machine.GP7)
	controls[3] = machine.Pin(machine.GP8)
	for _, control := range controls {
		control.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	pins[0] = machine.Pin(machine.GP11)
	pins[1] = machine.Pin(machine.GP9)
	pins[2] = machine.Pin(machine.GP13)
	pins[3] = machine.Pin(machine.GP15)
	pins[4] = machine.Pin(machine.GP16)
	pins[5] = machine.Pin(machine.GP10)
	pins[6] = machine.Pin(machine.GP12)
	pins[7] = machine.Pin(machine.GP14)

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
	machine.InitADC()
	adc0 := machine.ADC{Pin: machine.ADC0}
	adc1 := machine.ADC{Pin: machine.ADC1}
	adc2 := machine.ADC{Pin: machine.ADC2}

	adc0.Configure(machine.ADCConfig{})
	adc1.Configure(machine.ADCConfig{})
	adc2.Configure(machine.ADCConfig{})

	for {
		value := adc2.Get()
		value /= 7
		displayNumber(int(value))
	}
}
