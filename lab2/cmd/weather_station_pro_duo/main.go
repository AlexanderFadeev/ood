package main

import (
	"ood/lab2/display"
	"ood/lab2/stats_display"
	"ood/lab2/weather_data"
)

func main() {
	in := weather_data.WeatherData(weather_data.New())
	out := weather_data.New()

	d := display.New()
	sd := stats_display.New()

	in.Connect(weather_data.AllProBits, d.Display("In"), 1)
	out.ConnectPro(weather_data.AllProBits, d.DisplayPro("Out"), 1)
	in.Connect(weather_data.AllProBits, sd.DisplayStats("In"), 0)
	out.ConnectPro(weather_data.AllProBits, sd.DisplayStatsPro("Out"), 0)

	in.SetTemperature(1)
	in.SetPressure(2)
	in.SetHumidity(3)
	out.SetWind(100, 90)
	out.SetWind(50, 180)
	in.SetValues(3, 14, 15)
}
