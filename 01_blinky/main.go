package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/buzzer" // ← 追加
)

// ↓ 追加
type note struct {
	tone     float64
	duration float64
}

// ↑ 追加

func main() {
	//led := machine.LED
	led := machine.LCD_BACKLIGHT
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	cnt := 0

	// ↓ 追加
	bzrPin := machine.WIO_BUZZER
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	bzr := buzzer.New(bzrPin)

	notes := []note{
		{buzzer.C3, buzzer.Quarter},
		{buzzer.D3, buzzer.Quarter},
		{buzzer.E3, buzzer.Quarter},
	}
	// ↑ 追加

	for {
		cnt++
		fmt.Printf("cnt %d\r\n", cnt)
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)

		// ↓ 追加
		for _, n := range notes {
			bzr.Tone(n.tone, n.duration)
			time.Sleep(10 * time.Millisecond)
		}
		// ↑ 追加
		break
	}
}
