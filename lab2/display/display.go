package display

import (
	"fmt"

	"ood/lab2/weather_data"
)

type Displayer interface {
	Display(data weather_data.Getter)
	DisplayPro(data weather_data.GetterPro)
}

type displayer struct{}

func New() Displayer {
	return new(displayer)
}

func (d displayer) Display(data weather_data.Getter) {
	d.displayImpl(data)
	fmt.Println()
	fmt.Println()
}

func (d displayer) DisplayPro(data weather_data.GetterPro) {
	d.displayImpl(data)
	speed, dir := data.GetWind()
	fmt.Printf("Wind: %.1f m/s %.1f grad", speed, dir)
	fmt.Println()
	fmt.Println()
}

func (displayer) displayImpl(data weather_data.Getter) {
	fmt.Printf("Temp: %.1f, Press: %.1f, Hum: %.1f", data.GetTemperature(), data.GetPressure(), data.GetHumidity())

}
