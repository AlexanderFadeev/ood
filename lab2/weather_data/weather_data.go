package weather_data

import (
	"github.com/AlexanderFadeev/ood/lab2/signal"
)

type Setter interface {
	SetTemperature(float64)
	SetPressure(float64)
	SetHumidity(float64)
	SetValues(temperature, pressure, humidity float64)
}

type SetterPro interface {
	Setter
	SetWind(speed, direction float64)
	SetValuesPro(temperature, pressure, humidity, speed, direction float64)
}

type Getter interface {
	GetTemperature() float64
	GetPressure() float64
	GetHumidity() float64
}

type GetterPro interface {
	Getter
	GetWind() (speed, direction float64)
}

type Signal interface {
	DoOnTemperatureChange(slot FloatSlot, priority uint) signal.Connection
	DoOnPressureChange(slot FloatSlot, priority uint) signal.Connection
	DoOnHumidityChange(slot FloatSlot, priority uint) signal.Connection
}

type SignalPro interface {
	Signal

	DoOnWindChange(slot WindSlot, priority uint) signal.Connection
}

type WeatherData interface {
	Setter
	Getter
	Signal
}

type WeatherDataPro interface {
	SetterPro
	GetterPro
	SignalPro
}

type weatherData struct {
	temperature   float64
	pressure      float64
	humidity      float64
	windSpeed     float64
	windDirection float64

	onTemperatureChange FloatSignal
	onPressureChange    FloatSignal
	onHumidityChange    FloatSignal
	onWindChange        WindSignal
}

func New() WeatherDataPro {
	return &weatherData{
		onTemperatureChange: newFloatSignalAdapter(),
		onPressureChange:    newFloatSignalAdapter(),
		onHumidityChange:    newFloatSignalAdapter(),
		onWindChange:        newWindSignalAdapter(),
	}
}

func (wd *weatherData) SetTemperature(value float64) {
	wd.temperature = value
	wd.onTemperatureChange.Emit(value)
}

func (wd *weatherData) SetPressure(value float64) {
	wd.pressure = value
	wd.onPressureChange.Emit(value)
}

func (wd *weatherData) SetHumidity(value float64) {
	wd.humidity = value
	wd.onHumidityChange.Emit(value)
}

func (wd *weatherData) SetWind(speed float64, direction float64) {
	wd.windSpeed = speed
	wd.windDirection = direction
	wd.onWindChange.Emit(WindInfo{speed, direction})
}

func (wd *weatherData) SetValues(temperature, pressure, humidity float64) {
	wd.SetTemperature(temperature)
	wd.SetPressure(pressure)
	wd.SetHumidity(humidity)
}

func (wd *weatherData) SetValuesPro(temperature, pressure, humidity, speed, direction float64) {
	wd.SetValues(temperature, pressure, humidity)
	wd.SetWind(speed, direction)
}

func (wd *weatherData) GetTemperature() float64 {
	return wd.temperature
}

func (wd *weatherData) GetPressure() float64 {
	return wd.pressure
}

func (wd *weatherData) GetHumidity() float64 {
	return wd.humidity
}

func (wd *weatherData) GetWind() (float64, float64) {
	return wd.windSpeed, wd.windDirection
}

func (wd *weatherData) DoOnTemperatureChange(slot FloatSlot, priority uint) signal.Connection {
	return wd.onTemperatureChange.Connect(slot, priority)
}

func (wd *weatherData) DoOnPressureChange(slot FloatSlot, priority uint) signal.Connection {
	return wd.onPressureChange.Connect(slot, priority)

}

func (wd *weatherData) DoOnHumidityChange(slot FloatSlot, priority uint) signal.Connection {
	return wd.onHumidityChange.Connect(slot, priority)
}

func (wd *weatherData) DoOnWindChange(slot WindSlot, priority uint) signal.Connection {
	return wd.onWindChange.Connect(slot, priority)
}
