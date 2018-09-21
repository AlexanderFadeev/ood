package display

import (
	"fmt"

	"ood/lab2/weather_data"
)

type Displayer interface {
	Display(data weather_data.Getter)
}

type displayer struct{}

func New() Displayer {
	return new(displayer)
}

func (d displayer) Display(data weather_data.Getter) {
	fmt.Printf("Temp: %.1f, Press: %.1f, Hum: %.1f\n", data.GetTemperature(), data.GetPressure(), data.GetHumidity())
	fmt.Println()
}
