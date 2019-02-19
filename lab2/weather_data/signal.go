package weather_data

import "github.com/AlexanderFadeev/ood/lab2/signal"

type FloatSlot func(data float64)

func (s FloatSlot) exec(ctx interface{}) error {
	s(ctx.(float64))
	return nil
}

type FloatSignal interface {
	Connect(slot FloatSlot, priority uint) signal.Connection
	Emit(data float64)
}

type floatSignalAdapter struct {
	signal.Signal
}

func newFloatSignalAdapter() FloatSignal {
	return &floatSignalAdapter{
		Signal: signal.New(),
	}
}

func (s *floatSignalAdapter) Connect(slot FloatSlot, priority uint) signal.Connection {
	return s.Signal.Connect(slot.exec, priority)
}

func (s *floatSignalAdapter) Emit(data float64) {
	s.Signal.Emit(data)
}

type WindInfo struct {
	Speed     float64
	Direction float64
}

type WindSlot func(data WindInfo)

func (s WindSlot) exec(ctx interface{}) error {
	s(ctx.(WindInfo))
	return nil
}

type WindSignal interface {
	Connect(slot WindSlot, priority uint) signal.Connection
	Emit(data WindInfo)
}

type windSignalAdapter struct {
	signal.Signal
}

func newWindSignalAdapter() WindSignal {
	return &windSignalAdapter{
		Signal: signal.New(),
	}
}

func (s *windSignalAdapter) Connect(slot WindSlot, priority uint) signal.Connection {
	return s.Signal.Connect(slot.exec, priority)
}

func (s *windSignalAdapter) Emit(data WindInfo) {
	s.Signal.Emit(data)
}
