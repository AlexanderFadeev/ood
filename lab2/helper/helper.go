package helper

import (
	"github.com/AlexanderFadeev/ood/lab2/weather_data"
)

func WrapToFloatSlot(fn func()) weather_data.FloatSlot {
	return func(_ float64) {
		fn()
	}
}

func WrapToWindSlot(fn func()) weather_data.WindSlot {
	return func(_ weather_data.WindInfo) {
		fn()
	}
}
