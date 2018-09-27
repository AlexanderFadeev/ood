package weather_data

import "ood/lab2/signal"

type Slot func(data Getter)
type SlotPro func(data GetterPro)

func (s Slot) exec(ctx interface{}) error {
	s(ctx.(Getter))
	return nil
}

func (s SlotPro) exec(ctx interface{}) error {
	s(ctx.(GetterPro))
	return nil
}

type Signal interface {
	Connect(slot Slot, priority uint) signal.Connection
	Emit(data Getter)
}

type SignalPro interface {
	Signal
	ConnectPro(slot SlotPro, priority uint) signal.Connection
}

type signalAdapter struct {
	signal.Signal
}

func newSignalAdapter() SignalPro {
	return &signalAdapter{
		Signal: signal.New(),
	}
}

func (s *signalAdapter) Connect(slot Slot, priority uint) signal.Connection {
	return s.Signal.Connect(slot.exec, priority)
}

func (s *signalAdapter) ConnectPro(slot SlotPro, priority uint) signal.Connection {
	return s.Signal.Connect(slot.exec, priority)
}

func (s *signalAdapter) Emit(data Getter) {
	s.Signal.Emit(data)
}
