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

	conn := wd.ConnectPro(d.DisplayPro, 5)

	wd.SetTemperature(1)

	statsConn := wd.ConnectPro(sd.DisplayStatsPro, 42)
	defer statsConn.Close()

	wd.SetPressure(2)
	wd.SetHumidity(3)
	wd.SetWind(100, 90)
	wd.SetWind(50, 180)

	conn.Close()

	wd.SetValues(3, 14, 15)
}
