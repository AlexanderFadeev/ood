package main

import (
	"github.com/AlexanderFadeev/ood/lab2/display"
	"github.com/AlexanderFadeev/ood/lab2/helper"
	"github.com/AlexanderFadeev/ood/lab2/stats_display"
	"github.com/AlexanderFadeev/ood/lab2/weather_data"
)

const location = "Yoshkar-Ola"

func main() {
	wd := weather_data.New()
	d := display.New(wd, location)
	sd := stats_display.New(wd, location)

	conn1 := wd.DoOnTemperatureChange(helper.WrapToFloatSlot(d.Display), 5)
	conn2 := wd.DoOnPressureChange(helper.WrapToFloatSlot(d.Display), 5)
	conn3 := wd.DoOnHumidityChange(helper.WrapToFloatSlot(d.Display), 5)

	wd.SetTemperature(1)

	wd.DoOnTemperatureChange(helper.WrapToFloatSlot(sd.DisplayStats), 42)
	wd.DoOnPressureChange(helper.WrapToFloatSlot(sd.DisplayStats), 42)
	wd.DoOnHumidityChange(helper.WrapToFloatSlot(sd.DisplayStats), 42)

	wd.SetPressure(2)
	wd.SetHumidity(3)

	conn1.Close()
	conn2.Close()
	conn3.Close()

	wd.SetValues(3, 14, 15)
}
