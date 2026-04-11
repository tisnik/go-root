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
	}
}
