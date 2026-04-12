// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Demonstrační příklad použitý ve článcích:
//    Využití TinyGo při programování Raspberry Pi Pico: od GPIO až k PWM
//    https://www.root.cz/clanky/vyuziti-tinygo-pri-programovani-raspberry-pi-pico-od-gpio-az-k-pwm/
//
//    Využití TinyGo při programování Raspberry Pi Pico (2. část)
//    https://www.root.cz/clanky/vyuziti-tinygo-pri-programovani-raspberry-pi-pico-2-cast/
//

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

		for range 32 {
			for range 8 {
				rowsData.Set(true)
				rowsClock.Low()
				time.Sleep(1 * time.Millisecond)
				rowsClock.High()
			}
			rowsLatch.High()
			time.Sleep(1 * time.Millisecond)
			rowsLatch.Low()

			colsClock.High()
			time.Sleep(1 * time.Millisecond)
			colsClock.Low()

			colsLatch.High()
			time.Sleep(1 * time.Millisecond)
			colsLatch.Low()
		}
		time.Sleep(10 * time.Millisecond)
	}
}
