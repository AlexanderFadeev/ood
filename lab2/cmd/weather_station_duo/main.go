package main

import (
	"github.com/AlexanderFadeev/ood/lab2/display"
	"github.com/AlexanderFadeev/ood/lab2/helper"
	"github.com/AlexanderFadeev/ood/lab2/stats_display"
	"github.com/AlexanderFadeev/ood/lab2/weather_data"
)

const location = "Murmansk"

func main() {
	in := weather_data.New()
	out := weather_data.New()

	dIn := display.New(in, location+"In")
	sdIn := stats_display.New(in, location+"In")

	dOut := display.New(out, location+"Out")
	sdOut := stats_display.New(out, location+"Out")

	in.DoOnTemperatureChange(helper.WrapToFloatSlot(dIn.Display), 1)
	in.DoOnPressureChange(helper.WrapToFloatSlot(dIn.Display), 1)
	in.DoOnHumidityChange(helper.WrapToFloatSlot(dIn.Display), 1)

	in.DoOnTemperatureChange(helper.WrapToFloatSlot(sdIn.DisplayStats), 2)
	in.DoOnPressureChange(helper.WrapToFloatSlot(sdIn.DisplayStats), 2)
	in.DoOnHumidityChange(helper.WrapToFloatSlot(sdIn.DisplayStats), 2)

	out.DoOnTemperatureChange(helper.WrapToFloatSlot(dOut.Display), 1)
	out.DoOnPressureChange(helper.WrapToFloatSlot(dOut.Display), 1)
	out.DoOnHumidityChange(helper.WrapToFloatSlot(dOut.Display), 1)

	out.DoOnTemperatureChange(helper.WrapToFloatSlot(sdOut.DisplayStats), 2)
	out.DoOnPressureChange(helper.WrapToFloatSlot(sdOut.DisplayStats), 2)
	out.DoOnHumidityChange(helper.WrapToFloatSlot(sdOut.DisplayStats), 2)

	in.SetTemperature(1)
	in.SetPressure(2)
	in.SetHumidity(3)
	out.SetTemperature(42)
	out.SetPressure(9000)
	in.SetValues(3, 14, 15)
}
