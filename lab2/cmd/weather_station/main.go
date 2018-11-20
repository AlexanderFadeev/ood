package main

import (
	"ood/lab2/display"
	"ood/lab2/stats_display"
	"ood/lab2/weather_data"
)

func main() {
	wd := weather_data.New()
	d := display.New()
	sd := stats_display.New()

	conn := wd.Connect(weather_data.AllBits, d.Display(""), 5)

	wd.SetTemperature(1)

	statsConn := wd.Connect(weather_data.AllBits, sd.DisplayStats(""), 42)
	defer statsConn.Close()

	wd.SetPressure(2)
	wd.SetHumidity(3)

	conn.Close()

	wd.SetValues(3, 14, 15)
}
