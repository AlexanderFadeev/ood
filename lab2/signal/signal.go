package signal

type Signal interface {
	Connect(slot Slot, priority uint) Connection
	Emit(ctx interface{}) error
}

type signal struct {
	connections             uintToConnectionPriorityMap
	connectionsToDisconnect map[uint]struct{}
	lastConnectionID        uint
}

func New() Signal {
	return &signal{
		connections:             makeUintToConnectionPriorityMap(),
		connectionsToDisconnect: make(map[uint]struct{}),
	}
}

func (s *signal) Connect(slot Slot, priority uint) Connection {
	id := s.findUnusedConnectionID()
	s.lastConnectionID = id

	conn := &connection{
		id:     id,
		signal: s,
		slot:   slot,
	}

	s.connections.set(id, conn, priority)
	return conn
}

func (s *signal) Emit(ctx interface{}) (err error) {
	s.disconnectMarkedSlots()
	defer s.disconnectMarkedSlots()

	s.connections.iterate(func(_ uint, conn *connection) {
		slotErr := conn.slot(ctx)
		if slotErr != nil {
			err = slotErr
		}
	})

	return
}

func (s *signal) findUnusedConnectionID() uint {
	id := s.lastConnectionID + 1
	for {
		_, ok := s.connections.get(id)
		if !ok {
			return id
		}
		id++
	}
}

func (s *signal) disconnect(id uint) {
	s.connectionsToDisconnect[id] = struct{}{}
}

func (s *signal) disconnectMarkedSlots() {
	for id := range s.connectionsToDisconnect {
		s.connections.delete(id)
		delete(s.connectionsToDisconnect, id)
	}
}
