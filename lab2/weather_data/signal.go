package weather_data

import "github.com/AlexanderFadeev/ood/lab2/signal"

const (
	TemperatureBit uint = 1 << iota
	HumidityBit
	PressureBit
	WindBit
	AllBits    = PressureBit<<1 - 1
	AllProBits = WindBit<<1 - 1
)

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
	Connect(bitmap uint, slot Slot, priority uint) signal.Connection
	Emit(bitmap uint, data Getter)
}

type SignalPro interface {
	Signal
	ConnectPro(bitmap uint, slot SlotPro, priority uint) signal.Connection
}

type signalAdapter struct {
	signal.Signal
}

func newSignalAdapter() SignalPro {
	return &signalAdapter{
		Signal: signal.New(),
	}
}

func (s *signalAdapter) Connect(bitmap uint, slot Slot, priority uint) signal.Connection {
	return s.Signal.Connect(bitmap, slot.exec, priority)
}

func (s *signalAdapter) ConnectPro(bitmap uint, slot SlotPro, priority uint) signal.Connection {
	return s.Signal.Connect(bitmap, slot.exec, priority)
}

func (s *signalAdapter) Emit(bitmap uint, data Getter) {
	s.Signal.Emit(bitmap, data)
}
