package main

import (
	"machine"
	"time"
)

var raster []string = []string{
	//
	"***** * *   * *   *             ",
	"  *   * **  *  * *              ",
	"  *   * * * *   *   ****   **** ",
	"  *   * *  **   *  *    * *    *",
	"  *   * *   *   *  *      *    *",
	"                   *   ** *    *",
	"                   *    * *    *",
	"                    ****   **** ",
}

func main() {
	machine.InitADC()
	adc0 := machine.ADC{Pin: machine.ADC0}
	adc1 := machine.ADC{Pin: machine.ADC1}
	adc2 := machine.ADC{Pin: machine.ADC2}

	adc0.Configure(machine.ADCConfig{})
	adc1.Configure(machine.ADCConfig{})
	adc2.Configure(machine.ADCConfig{})

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

		var scroll = int(adc2.Get()) >> 11

		for x := range 32 {
			for y := range 8 {
				pixel := raster[7-y][(x+scroll)%32]
				value := pixel != ' '
				rowsData.Set(value)

				rowsClock.Low()
				time.Sleep(1 * time.Microsecond)
				rowsClock.High()

				time.Sleep(10 * time.Microsecond)
			}
			rowsLatch.High()
			time.Sleep(1 * time.Microsecond)
			rowsLatch.Low()

			colsClock.High()
			time.Sleep(1 * time.Microsecond)
			colsClock.Low()

			colsLatch.High()
			time.Sleep(1 * time.Microsecond)
			colsLatch.Low()
		}
		time.Sleep(10 * time.Microsecond)
	}
}
