package stats_display

import (
	"fmt"

	"ood/lab2/weather_data"
)

type StatsDisplayer interface {
	DisplayStats(weather_data.Getter)
	DisplayStatsPro(weather_data.GetterPro)
}

type statsDisplayer struct {
	temperatureStatsCollector collector
	pressureStatsCollector    collector
	humidityStatsCollector    collector
	windStatsCollector        windCollector
}

func New() StatsDisplayer {
	return &statsDisplayer{
		temperatureStatsCollector: new(collectorImpl),
		pressureStatsCollector:    new(collectorImpl),
		humidityStatsCollector:    new(collectorImpl),
		windStatsCollector:        new(windCollectorImpl),
	}
}

func (sd *statsDisplayer) DisplayStats(data weather_data.Getter) {
	sd.update(data)
	sd.displayImpl(data)
	fmt.Println()
}

func (sd *statsDisplayer) DisplayStatsPro(data weather_data.GetterPro) {
	sd.updatePro(data)
	sd.displayImpl(data)
	sd.displayWindStats()
	fmt.Println()
}

func (sd *statsDisplayer) displayImpl(data weather_data.Getter) {
	sd.displayCollectorStats("Temp", sd.temperatureStatsCollector)
	sd.displayCollectorStats("Press", sd.pressureStatsCollector)
	sd.displayCollectorStats("Hum", sd.humidityStatsCollector)
}

func (sd statsDisplayer) displayCollectorStats(name string, collector collector) {
	fmt.Printf("Min%s: %.1f, Max%s: %.1f, Avg%s: %.1f\n",
		name, collector.GetMin(),
		name, collector.GetMax(),
		name, collector.GetAverage(),
	)
}

func (sd *statsDisplayer) displayWindStats() {
	avgSpeed, avgDir := sd.windStatsCollector.GetAverage()
	fmt.Printf("MaxWind: %.1f, MinWind: %.1f, AvgWind: %.1f, AvgWindDirection: %.1f",
		sd.windStatsCollector.GetMaxSpeed(), sd.windStatsCollector.GetMinSpeed(), avgSpeed, avgDir)
}

func (sd *statsDisplayer) updatePro(data weather_data.GetterPro) {
	sd.update(data)
	sd.windStatsCollector.AddWindValue(data.GetWind())
}

func (sd *statsDisplayer) update(data weather_data.Getter) {
	sd.temperatureStatsCollector.AddValue(data.GetTemperature())
	sd.pressureStatsCollector.AddValue(data.GetPressure())
	sd.humidityStatsCollector.AddValue(data.GetHumidity())
}
