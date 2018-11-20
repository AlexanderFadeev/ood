package signal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const bitmap = 1

func TestSignalCallSlot(t *testing.T) {
	signal := New()
	var c value

	signal.Connect(bitmap, c.set, 0)

	assert.Equal(t, nil, c.get())
	signal.Emit(bitmap, 42)
	assert.Equal(t, 42, c.get())
}

func TestSignalCallMultipleSlots(t *testing.T) {
	signal := New()
	var c1 value
	var c2 value

	signal.Connect(bitmap, c1.set, 0)
	signal.Connect(bitmap, c2.set, 0)

	signal.Emit(bitmap, 42)
	assert.Equal(t, 42, c1.get())
	assert.Equal(t, 42, c2.get())
}

func TestSignalEmittedValue(t *testing.T) {
	signal := New()
	var c value

	signal.Connect(bitmap, c.set, 0)

	assert.Equal(t, nil, c.get())
	signal.Emit(bitmap, 42)
	assert.Equal(t, 42, c.get())
}

func TestSignalDisconnect(t *testing.T) {
	signal := New()
	var c value

	conn := signal.Connect(bitmap, c.set, 0)

	signal.Emit(bitmap, 42)
	assert.Equal(t, 42, c.get())

	conn.Close()

	signal.Emit(bitmap, nil)
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
	conn := signal.Connect(bitmap, noOpSlot, 0)
	signal.Connect(bitmap, makeCloseConnectionSlot(conn), 42)
	signal.Emit(bitmap, nil)
}

func TestSignalDisconnectWhileEmittingAfterCalling(t *testing.T) {
	signal := New()
	conn := signal.Connect(bitmap, noOpSlot, 42)
	signal.Connect(bitmap, makeCloseConnectionSlot(conn), 0)
	signal.Emit(bitmap, nil)
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

	signal.Connect(bitmap, makeSqrSlot(&i), 1)
	signal.Connect(bitmap, makeIncSlot(&i), 3)
	signal.Connect(bitmap, makeDoubleSlot((&i)), 2)

	signal.Emit(bitmap, nil)
	assert.Equal(t, 4, i)
}

func TestSignalBitmap(t *testing.T) {
	signal := New()
	i := 0

	signal.Connect(1, makeIncSlot(&i), 1)
	signal.Connect(2, makeSqrSlot(&i), 2)

	signal.Emit(1, nil) // 0+1
	signal.Emit(1, nil) // 1+1
	signal.Emit(2, nil) // 2*2
	signal.Emit(3, nil) // 4*4+1

	assert.Equal(t, 17, i)
}
