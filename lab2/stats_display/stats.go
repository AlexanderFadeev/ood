package stats_display

type collector interface {
	AddValue(float64)
	GetAverage() float64
	GetMax() float64
	GetMin() float64
}

type collectorImpl struct {
	sum   float64
	max   float64
	min   float64
	count uint
}

func (c *collectorImpl) AddValue(value float64) {
	if c.count == 0 || value > c.max {
		c.max = value
	}

	if c.count == 0 || value < c.min {
		c.min = value
	}

	c.sum += value
	c.count++
}

func (c *collectorImpl) GetAverage() float64 {
	if c.count == 0 {
		return 0
	}

	return c.sum / float64(c.count)
}

func (c *collectorImpl) GetMax() float64 {
	return c.max
}

func (c *collectorImpl) GetMin() float64 {
	return c.min
}
