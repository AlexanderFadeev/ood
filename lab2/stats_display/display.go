package stats_display

import (
	"fmt"

	"github.com/AlexanderFadeev/ood/lab2/weather_data"
)

type StatsDisplayer interface {
	DisplayStats(location string) weather_data.Slot
	DisplayStatsPro(location string) weather_data.SlotPro
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

func (sd *statsDisplayer) DisplayStats(location string) weather_data.Slot {
	return func(data weather_data.Getter) {
		sd.update(data)
		sd.displayImpl(location, data)
		fmt.Println()
	}
}

func (sd *statsDisplayer) DisplayStatsPro(location string) weather_data.SlotPro {
	return func(data weather_data.GetterPro) {
		sd.updatePro(data)
		sd.displayImpl(location, data)
		sd.displayWindStats()
		fmt.Println()
	}
}

func (sd *statsDisplayer) displayImpl(location string, data weather_data.Getter) {
	fmt.Printf("#%s ", location)
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
