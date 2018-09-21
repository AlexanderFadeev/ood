package weather_data

type Setter interface {
	SetTemperature(float64)
	SetPressure(float64)
	SetHumidity(float64)
	SetValues(temperature float64, pressure float64, humidity float64)
}

type Getter interface {
	GetTemperature() float64
	GetPressure() float64
	GetHumidity() float64
}

type WeatherData interface {
	Setter
	Getter
	Signal
}

type weatherData struct {
	temperature float64
	pressure    float64
	humidity    float64
	Signal
}

func New() WeatherData {
	return &weatherData{
		Signal: newSignalAdapter(),
	}
}

func (wd *weatherData) SetTemperature(value float64) {
	wd.temperature = value
	wd.notifyObservers()
}

func (wd *weatherData) SetPressure(value float64) {
	wd.pressure = value
	wd.notifyObservers()
}

func (wd *weatherData) SetHumidity(value float64) {
	wd.humidity = value
	wd.notifyObservers()
}

func (wd *weatherData) SetValues(temperature float64, pressure float64, humidity float64) {
	wd.temperature = temperature
	wd.pressure = pressure
	wd.humidity = humidity
	wd.notifyObservers()
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

func (wd *weatherData) notifyObservers() {
	wd.Signal.Emit(wd.getSelfCopyPtr())
}

func (wd *weatherData) getSelfCopyPtr() *weatherData {
	wdCopy := *wd
	return &wdCopy
}
