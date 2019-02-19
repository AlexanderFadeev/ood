package display

import (
	"fmt"

	"github.com/AlexanderFadeev/ood/lab2/weather_data"
)

type Displayer interface {
	Display(location string) weather_data.Slot
	DisplayPro(location string) weather_data.SlotPro
}

type displayer struct{}

func New() Displayer {
	return new(displayer)
}

func (d displayer) Display(location string) weather_data.Slot {
	return func(data weather_data.Getter) {
		d.displayImpl(location, data)
		fmt.Println()
	}
}

func (d displayer) DisplayPro(location string) weather_data.SlotPro {
	return func(data weather_data.GetterPro) {
		d.displayImpl(location, data)
		speed, dir := data.GetWind()
		fmt.Printf(" Wind: %.1f m/s %.1f grad", speed, dir)
		fmt.Println()
	}
}

func (displayer) displayImpl(location string, data weather_data.Getter) {
	fmt.Printf("#%s Temp: %.1f, Press: %.1f, Hum: %.1f",
		location, data.GetTemperature(), data.GetPressure(), data.GetHumidity())
}
