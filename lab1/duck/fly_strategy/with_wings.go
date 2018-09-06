package fly_strategy

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

func WithWings() FlyStrategy {
	flightsCount := 0

	return func() {
		flightsCount++
		fmt.Printf("I'm flying with wings %s time!!!\n", humanize.Ordinal(flightsCount))
	}
}
