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
	pwm0 := machine.PWM0
	pwmPin := machine.GP0
	pwmPin.Configure(machine.PinConfig{Mode: machine.PinPWM})

	led := machine.GP2
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.Low()

	err := pwm0.Configure(machine.PWMConfig{
		Period: 1e9 / 500,
	})
	if err != nil {
		return
	}

	channel, err := pwm0.Channel(pwmPin)
	if err != nil {
		return
	}
	led.High()

	for {
		for divider := 1; divider <= 10; divider++ {
			pwm0.Set(channel, pwm0.Top()/uint32(divider))
			time.Sleep(100 * time.Millisecond)
		}
		for divider := 10; divider >= 1; divider-- {
			pwm0.Set(channel, pwm0.Top()/uint32(divider))
			time.Sleep(100 * time.Millisecond)
		}
	}
}
