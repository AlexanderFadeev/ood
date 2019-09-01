package main

import (
	"github.com/AlexanderFadeev/ood/lab2/display"
	"github.com/AlexanderFadeev/ood/lab2/helper"
	"github.com/AlexanderFadeev/ood/lab2/stats_display"
	"github.com/AlexanderFadeev/ood/lab2/weather_data"
)

const location = "Lake Baikal"

func main() {
	wd := weather_data.New()

	d := display.New(wd, location)
	sd := stats_display.New(wd, location)

	wd.DoOnTemperatureChange(helper.WrapToFloatSlot(d.DisplayPro), 1)
	wd.DoOnPressureChange(helper.WrapToFloatSlot(d.DisplayPro), 1)

	wd.DoOnTemperatureChange(helper.WrapToFloatSlot(sd.DisplayStatsPro), 0)
	wd.DoOnPressureChange(helper.WrapToFloatSlot(sd.DisplayStatsPro), 0)

	wd.SetTemperature(1)
	wd.SetPressure(2)
	wd.SetHumidity(3)
	wd.SetWind(100, 90)
	wd.SetWind(50, 180)
	wd.SetValues(3, 14, 15)
}
