package signal

type Signal interface {
	Connect(slot Slot, priority uint) Connection
	Emit(ctx interface{}) error
}

type signal struct {
	slots             uintToSlotPriorityMap
	slotsToDisconnect map[uint]struct{}
	lastConnectionID  uint
}

func New() Signal {
	return &signal{
		slots:             makeUintToSlotPriorityMap(),
		slotsToDisconnect: make(map[uint]struct{}),
	}
}

func (s *signal) Connect(slot Slot, priority uint) Connection {
	id := s.findUnusedConnectionID()
	s.slots.set(id, slot, priority)

	s.lastConnectionID = id
	return &connection{
		id:     id,
		signal: s,
	}
}

func (s *signal) Emit(ctx interface{}) (err error) {
	s.disconnectMarkedSlots()
	defer s.disconnectMarkedSlots()

	s.slots.iterate(func(_ uint, slot Slot) {
		slotErr := slot(ctx)
		if slotErr != nil {
			err = slotErr
		}
	})

	return
}

func (s *signal) findUnusedConnectionID() uint {
	id := s.lastConnectionID + 1
	for {
		_, ok := s.slots.get(id)
		if !ok {
			return id
		}
		id++
	}
}

func (s *signal) disconnect(id uint) {
	s.slotsToDisconnect[id] = struct{}{}
}

func (s *signal) disconnectMarkedSlots() {
	for id := range s.slotsToDisconnect {
		s.slots.delete(id)
		delete(s.slotsToDisconnect, id)
	}
}
