package signal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignalCallSlot(t *testing.T) {
	signal := New()
	var c value

	signal.Connect(c.set, 0)

	assert.Equal(t, nil, c.get())
	signal.Emit(42)
	assert.Equal(t, 42, c.get())
}

func TestSignalCallMultipleSlots(t *testing.T) {
	signal := New()
	var c1 value
	var c2 value

	signal.Connect(c1.set, 0)
	signal.Connect(c2.set, 0)

	signal.Emit(42)
	assert.Equal(t, 42, c1.get())
	assert.Equal(t, 42, c2.get())
}

func TestSignalEmittedValue(t *testing.T) {
	signal := New()
	var c value

	signal.Connect(c.set, 0)

	assert.Equal(t, nil, c.get())
	signal.Emit(42)
	assert.Equal(t, 42, c.get())
}

func TestSignalDisconnect(t *testing.T) {
	signal := New()
	var c value

	conn := signal.Connect(c.set, 0)

	signal.Emit(42)
	assert.Equal(t, 42, c.get())

	conn.Close()

	signal.Emit(nil)
	assert.Equal(t, 42, c.get())
}

func noOpSlot(_ interface{}) error {
	return nil
}

func makeCloseConnectionSlot(conn Connection) Slot {
	return func(interface{}) error {
		conn.Close()
		return nil
	}
}

func TestSignalDisconnectWhileEmittingBeforeCalling(t *testing.T) {
	signal := New()
	conn := signal.Connect(noOpSlot, 0)
	signal.Connect(makeCloseConnectionSlot(conn), 42)
	signal.Emit(nil)
}

func TestSignalDisconnectWhileEmittingAfterCalling(t *testing.T) {
	signal := New()
	conn := signal.Connect(noOpSlot, 42)
	signal.Connect(makeCloseConnectionSlot(conn), 0)
	signal.Emit(nil)
}

func makeIncSlot(i *int) Slot {
	return func(_ interface{}) error {
		*i++
		return nil
	}
}

func makeDoubleSlot(i *int) Slot {
	return func(_ interface{}) error {
		*i *= 2
		return nil
	}
}

func makeSqrSlot(i *int) Slot {
	return func(_ interface{}) error {
		*i *= *i
		return nil
	}
}

func TestSignalPriority(t *testing.T) {
	signal := New()
	i := 0

	signal.Connect(makeSqrSlot(&i), 1)
	signal.Connect(makeIncSlot(&i), 3)
	signal.Connect(makeDoubleSlot(&i), 2)

	signal.Emit(nil)
	assert.Equal(t, 4, i)
}
