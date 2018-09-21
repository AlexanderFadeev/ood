package weather_data

import "ood/lab2/signal"

type Slot func(data Getter)

func (s Slot) exec(ctx interface{}) error {
	s(ctx.(Getter))
	return nil
}

type Signal interface {
	Connect(slot Slot, priority uint) signal.Connection
	Emit(data Getter)
}

type signalAdapter struct {
	signal.Signal
}

func newSignalAdapter() Signal {
	return &signalAdapter{
		Signal: signal.New(),
	}
}

func (s *signalAdapter) Connect(slot Slot, priority uint) signal.Connection {
	return s.Signal.Connect(slot.exec, priority)
}

func (s *signalAdapter) Emit(data Getter) {
	s.Signal.Emit(data)
}
