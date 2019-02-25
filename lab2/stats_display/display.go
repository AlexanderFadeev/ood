package stats_display

import (
	"fmt"

	"github.com/AlexanderFadeev/ood/lab2/weather_data"
)

type StatsDisplayer interface {
	DisplayStats()
	DisplayStatsPro()
}

type statsDisplayer struct {
	wd       weather_data.GetterPro
	location string

	temperatureStatsCollector collector
	pressureStatsCollector    collector
	humidityStatsCollector    collector
	windStatsCollector        windCollector
}

func New(wd weather_data.GetterPro, location string) StatsDisplayer {
	return &statsDisplayer{
		wd:       wd,
		location: location,

		temperatureStatsCollector: new(collectorImpl),
		pressureStatsCollector:    new(collectorImpl),
		humidityStatsCollector:    new(collectorImpl),
		windStatsCollector:        new(windCollectorImpl),
	}
}

func (sd *statsDisplayer) DisplayStats() {
	sd.update()
	sd.displayImpl()
	fmt.Println()
}

func (sd *statsDisplayer) DisplayStatsPro() {
	sd.updatePro()
	sd.displayImpl()
	sd.displayWindStats()
	fmt.Println()
}

func (sd *statsDisplayer) displayImpl() {
	fmt.Printf("#%s ", sd.location)
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

func (sd *statsDisplayer) updatePro() {
	sd.update()
	sd.windStatsCollector.AddWindValue(sd.wd.GetWind())
}

func (sd *statsDisplayer) update() {
	sd.temperatureStatsCollector.AddValue(sd.wd.GetTemperature())
	sd.pressureStatsCollector.AddValue(sd.wd.GetPressure())
	sd.humidityStatsCollector.AddValue(sd.wd.GetHumidity())
}
