package stats_display

import (
	"fmt"

	"ood/lab2/weather_data"
)

type StatsDisplayer interface {
	DisplayStats(weather_data.Getter)
}

type statsDisplayer struct {
	temperatureStatsCollector collector
	pressureStatsCollector    collector
	humidityStatsCollector    collector
}

func New() StatsDisplayer {
	return &statsDisplayer{
		temperatureStatsCollector: new(collectorImpl),
		pressureStatsCollector:    new(collectorImpl),
		humidityStatsCollector:    new(collectorImpl),
	}
}

func (sd *statsDisplayer) DisplayStats(data weather_data.Getter) {
	sd.update(data)
	sd.displayCollectorStats("Temp", sd.temperatureStatsCollector)
	sd.displayCollectorStats("Press", sd.pressureStatsCollector)
	sd.displayCollectorStats("Hum", sd.humidityStatsCollector)
	fmt.Println()
}

func (sd statsDisplayer) displayCollectorStats(name string, collector collector) {
	fmt.Printf("Min%s: %.1f, Max%s: %.1f, Avg%s: %.1f\n",
		name, collector.GetMin(),
		name, collector.GetMax(),
		name, collector.GetAverage(),
	)
}

func (sd *statsDisplayer) update(data weather_data.Getter) {
	sd.temperatureStatsCollector.AddValue(data.GetTemperature())
	sd.pressureStatsCollector.AddValue(data.GetPressure())
	sd.humidityStatsCollector.AddValue(data.GetHumidity())
}
