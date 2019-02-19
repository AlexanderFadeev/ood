package main

import (
	"github.com/AlexanderFadeev/ood/lab2/display"
	"github.com/AlexanderFadeev/ood/lab2/stats_display"
	"github.com/AlexanderFadeev/ood/lab2/weather_data"
)

const fishingEventsBits = weather_data.PressureBit | weather_data.TemperatureBit

func main() {
	wd := weather_data.New()

	d := display.New()
	sd := stats_display.New()

	wd.ConnectPro(fishingEventsBits, d.DisplayPro("Out"), 1)
	wd.ConnectPro(fishingEventsBits, sd.DisplayStatsPro("Out"), 0)

	wd.SetTemperature(1)
	wd.SetPressure(2)
	wd.SetHumidity(3)
	wd.SetWind(100, 90)
	wd.SetWind(50, 180)
	wd.SetValues(3, 14, 15)
}
