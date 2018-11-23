package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	raspberry := raspi.NewAdaptor()
	oled := i2c.NewSSD1306Driver(raspberry)

	state := 0

	work := func() {
		gobot.Every(time.Second, func() {
			state = (state + 1) % 2
			// log.Println("state =", state)
			oled.Clear()
			oled.Set(10, 10, state)
			oled.Display()
		})
	}

	robot := gobot.NewRobot(
		"oled-rest",
		[]gobot.Connection{raspberry},
		[]gobot.Device{oled},
		work)
	robot.Start()
}
