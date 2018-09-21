package signal

type Connection interface {
	Close()
}

type connection struct {
	id     uint
	signal *signal
}

func (c *connection) Close() {
	c.signal.disconnect(c.id)
}
