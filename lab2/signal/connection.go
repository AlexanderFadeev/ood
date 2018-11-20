package signal

type Connection interface {
	Close()
}

type connection struct {
	id     uint
	bitmap uint
	signal *signal
	slot   Slot
}

func (c *connection) Close() {
	c.signal.disconnect(c.id)
}
