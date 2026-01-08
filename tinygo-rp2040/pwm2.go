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
