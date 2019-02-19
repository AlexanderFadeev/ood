package display

import (
	"fmt"

	"github.com/AlexanderFadeev/ood/lab2/weather_data"
)

type Displayer interface {
	Display()
	DisplayPro()
}

type displayer struct {
	wd       weather_data.GetterPro
	location string
}

func New(wd weather_data.WeatherDataPro, location string) Displayer {
	return &displayer{
		wd:       wd,
		location: location,
	}
}

func (d displayer) Display() {
	d.displayImpl()
	fmt.Println()
}

func (d displayer) DisplayPro() {
	d.displayImpl()
	speed, dir := d.wd.GetWind()
	fmt.Printf(" Wind: %.1f m/s %.1f grad", speed, dir)
	fmt.Println()
}

func (d displayer) displayImpl() {
	fmt.Printf("#%s Temp: %.1f, Press: %.1f, Hum: %.1f",
		d.location, d.wd.GetTemperature(), d.wd.GetPressure(), d.wd.GetHumidity())
}
