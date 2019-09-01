package main

import (
	"github.com/AlexanderFadeev/ood/lab2/display"
	"github.com/AlexanderFadeev/ood/lab2/helper"
	"github.com/AlexanderFadeev/ood/lab2/stats_display"
	"github.com/AlexanderFadeev/ood/lab2/weather_data"
)

const location = "North Pole"

func main() {
	in := weather_data.New()
	out := weather_data.New()

	dIn := display.New(in, location+"In")
	sdIn := stats_display.New(in, location+"In")

	dOut := display.New(out, location+"Out")
	sdOut := stats_display.New(out, location+"Out")

	in.DoOnTemperatureChange(helper.WrapToFloatSlot(dIn.DisplayPro), 1)
	in.DoOnPressureChange(helper.WrapToFloatSlot(dIn.DisplayPro), 1)
	in.DoOnHumidityChange(helper.WrapToFloatSlot(dIn.DisplayPro), 1)
	in.DoOnWindChange(helper.WrapToWindSlot(dIn.DisplayPro), 1)

	in.DoOnTemperatureChange(helper.WrapToFloatSlot(sdIn.DisplayStatsPro), 2)
	in.DoOnPressureChange(helper.WrapToFloatSlot(sdIn.DisplayStatsPro), 2)
	in.DoOnHumidityChange(helper.WrapToFloatSlot(sdIn.DisplayStatsPro), 2)
	in.DoOnWindChange(helper.WrapToWindSlot(sdIn.DisplayStatsPro), 2)

	out.DoOnTemperatureChange(helper.WrapToFloatSlot(dOut.DisplayPro), 1)
	out.DoOnPressureChange(helper.WrapToFloatSlot(dOut.DisplayPro), 1)
	out.DoOnHumidityChange(helper.WrapToFloatSlot(dOut.DisplayPro), 1)
	out.DoOnWindChange(helper.WrapToWindSlot(dOut.DisplayPro), 1)

	out.DoOnTemperatureChange(helper.WrapToFloatSlot(sdOut.DisplayStatsPro), 2)
	out.DoOnPressureChange(helper.WrapToFloatSlot(sdOut.DisplayStatsPro), 2)
	out.DoOnHumidityChange(helper.WrapToFloatSlot(sdOut.DisplayStatsPro), 2)
	out.DoOnWindChange(helper.WrapToWindSlot(sdOut.DisplayStatsPro), 2)

	in.SetTemperature(1)
	in.SetPressure(2)
	in.SetHumidity(3)
	out.SetWind(100, 90)
	out.SetWind(50, 180)
	in.SetValues(3, 14, 15)
}
