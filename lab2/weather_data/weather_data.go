package weather_data

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
	SignalPro
}

func New() WeatherDataPro {
	return &weatherData{
		SignalPro: newSignalAdapter(),
	}
}

func (wd *weatherData) SetTemperature(value float64) {
	wd.temperature = value
	wd.notifyObservers(TemperatureBit)
}

func (wd *weatherData) SetPressure(value float64) {
	wd.pressure = value
	wd.notifyObservers(PressureBit)
}

func (wd *weatherData) SetHumidity(value float64) {
	wd.humidity = value
	wd.notifyObservers(HumidityBit)
}

func (wd *weatherData) SetWind(speed float64, direction float64) {
	wd.windSpeed = speed
	wd.windDirection = direction
	wd.notifyObservers(WindBit)
}

func (wd *weatherData) SetValues(temperature, pressure, humidity float64) {
	wd.temperature = temperature
	wd.pressure = pressure
	wd.humidity = humidity
	wd.notifyObservers(AllBits)
}

func (wd *weatherData) SetValuesPro(temperature, pressure, humidity, speed, direction float64) {
	wd.temperature = temperature
	wd.pressure = pressure
	wd.humidity = humidity
	wd.windSpeed = speed
	wd.windDirection = direction
	wd.notifyObservers(AllProBits)
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

func (wd *weatherData) notifyObservers(bitmap uint) {
	wd.SignalPro.Emit(bitmap, wd.getSelfCopyPtr())
}

func (wd *weatherData) getSelfCopyPtr() *weatherData {
	wdCopy := *wd
	return &wdCopy
}
