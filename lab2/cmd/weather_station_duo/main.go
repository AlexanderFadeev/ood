package main

import (
	"ood/lab2/display"
	"ood/lab2/stats_display"
	"ood/lab2/weather_data"
)

func main() {
	in := weather_data.New()
	out := weather_data.New()

	d := display.New()
	sd := stats_display.New()

	in.ConnectPro(d.DisplayPro("In"), 1)
	out.ConnectPro(d.DisplayPro("Out"), 1)
	in.ConnectPro(sd.DisplayStatsPro("In"), 0)
	out.ConnectPro(sd.DisplayStatsPro("Out"), 0)

	in.SetTemperature(1)
	in.SetPressure(2)
	in.SetHumidity(3)
	out.SetWind(100, 90)
	out.SetWind(50, 180)
	in.SetValues(3, 14, 15)
}