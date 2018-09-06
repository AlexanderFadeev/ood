package fly_strategy

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

type WithWings struct {
	flightsCount int
}

func (w *WithWings) Fly() {
	w.flightsCount++
	fmt.Printf("I'm flying with wings %s time!!!\n", humanize.Ordinal(w.flightsCount))
}
