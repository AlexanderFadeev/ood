package stats_display

import "math"

type windCollector interface {
	AddWindValue(speed, direction float64)
	GetAverage() (speed, direction float64)
	GetMaxSpeed() float64
	GetMinSpeed() float64
}

type windCollectorImpl struct {
	speedCollector collectorImpl
	sinAccum       float64
	cosAccum       float64
	count          uint
}

func (wc *windCollectorImpl) AddWindValue(speed, direction float64) {
	wc.speedCollector.AddValue(speed)
	wc.sinAccum += speed * math.Sin(wc.degToRad(direction))
	wc.cosAccum += speed * math.Cos(wc.degToRad(direction))
	wc.count++
}

func (wc *windCollectorImpl) GetAverage() (float64, float64) {
	avgSpeed := math.Hypot(wc.sinAccum, wc.cosAccum) / float64(wc.count)
	avgDirectionRad := math.Atan2(wc.sinAccum, wc.cosAccum)
	return avgSpeed, wc.radToNormalizedDeg(avgDirectionRad)
}

func (wc *windCollectorImpl) GetMaxSpeed() float64 {
	return wc.speedCollector.GetMax()
}

func (wc *windCollectorImpl) GetMinSpeed() float64 {
	return wc.speedCollector.GetMin()
}

func (windCollectorImpl) degToRad(deg float64) float64 {
	return deg / 180 * math.Pi
}

func (windCollectorImpl) radToNormalizedDeg(rad float64) float64 {
	deg := rad / math.Pi * 180
	if deg < 0 {
		deg += 360
	}
	return deg
}
