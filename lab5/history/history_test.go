package history

import (
	"testing"

	"github.com/AlexanderFadeev/ood/lab5/command"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const sampleSize = 1

var sampleError = errors.New("Sample error")

func TestRecord(t *testing.T) {
	cmd := new(command.MockCommand)
	cmd.On("Execute").Return(nil)

	history := New(sampleSize)
	err := history.Record(cmd)
	assert.Nil(t, err)

	cmd.AssertNumberOfCalls(t, "Execute", 1)
}

func TestRecordError(t *testing.T) {
	cmd := new(command.MockCommand)
	cmd.On("Execute").Return(sampleError)

	history := New(sampleSize)
	err := history.Record(cmd)
	assert.NotNil(t, err)

	cmd.AssertNumberOfCalls(t, "Execute", 1)
}

func TestUndo(t *testing.T) {
	cmd := new(command.MockCommand)
	cmd.On("Execute").Return(nil)
	cmd.On("Unexecute").Return(nil)

	history := New(sampleSize)
	assert.False(t, history.CanUndo())
	assert.Panics(t, func() { history.Undo() })

	history.Record(cmd)
	assert.True(t, history.CanUndo())

	history.Undo()
	assert.False(t, history.CanUndo())
	assert.Panics(t, func() { history.Undo() })

	cmd.AssertNumberOfCalls(t, "Execute", 1)
	cmd.AssertNumberOfCalls(t, "Unexecute", 1)
}

func TestUndoError(t *testing.T) {
	cmd := new(command.MockCommand)
	cmd.On("Execute").Return(nil)
	cmd.On("Unexecute").Return(sampleError)

	history := New(sampleSize)
	history.Record(cmd)
	err := history.Undo()
	assert.NotNil(t, err)

	cmd.AssertNumberOfCalls(t, "Execute", 1)
	cmd.AssertNumberOfCalls(t, "Unexecute", 1)
}

func TestRedo(t *testing.T) {
	cmd := new(command.MockCommand)
	cmd.On("Execute").Return(nil)
	cmd.On("Unexecute").Return(nil)

	history := New(sampleSize)
	assert.False(t, history.CanRedo())
	assert.Panics(t, func() { history.Redo() })

	history.Record(cmd)
	assert.False(t, history.CanRedo())
	assert.Panics(t, func() { history.Redo() })

	history.Undo()
	assert.True(t, history.CanRedo())

	history.Redo()
	assert.False(t, history.CanRedo())
	assert.Panics(t, func() { history.Redo() })

	cmd.AssertNumberOfCalls(t, "Execute", 2)
	cmd.AssertNumberOfCalls(t, "Unexecute", 1)
}

func TestRedoError(t *testing.T) {
	cmd := new(command.MockCommand)
	calledTime := 0

	cmd.On("Execute").Return(func() error {
		calledTime++
		if calledTime == 2 {
			return sampleError
		}
		return nil
	})
	cmd.On("Unexecute").Return(nil)

	history := New(sampleSize)
	history.Record(cmd)
	history.Undo()
	err := history.Redo()
	assert.NotNil(t, err)

	cmd.AssertNumberOfCalls(t, "Execute", 2)
	cmd.AssertNumberOfCalls(t, "Unexecute", 1)
}

func TestRecordAfterUndo(t *testing.T) {
	cmdA := new(command.MockCommand)
	cmdA.On("Execute").Return(nil)
	cmdA.On("Unexecute").Return(nil)
	cmdA.On("Release").Return(nil)
	cmdB := new(command.MockCommand)
	cmdB.On("Execute").Return(nil)

	history := New(sampleSize)
	history.Record(cmdA)
	history.Undo()
	history.Record(cmdB)

	cmdA.AssertNumberOfCalls(t, "Execute", 1)
	cmdA.AssertNumberOfCalls(t, "Unexecute", 1)
	cmdA.AssertNumberOfCalls(t, "Release", 1)
	cmdB.AssertNumberOfCalls(t, "Execute", 1)
}

func TestReleaseError(t *testing.T) {
	cmdA := new(command.MockCommand)
	cmdA.On("Execute").Return(nil)
	cmdA.On("Unexecute").Return(nil)
	cmdA.On("Release").Return(sampleError)
	cmdB := new(command.MockCommand)
	cmdB.On("Execute").Return(nil)

	history := New(sampleSize)
	history.Record(cmdA)
	history.Undo()
	err := history.Record(cmdB)
	assert.NotNil(t, err)

	cmdA.AssertNumberOfCalls(t, "Execute", 1)
	cmdA.AssertNumberOfCalls(t, "Unexecute", 1)
	cmdA.AssertNumberOfCalls(t, "Release", 1)
	cmdB.AssertNumberOfCalls(t, "Execute", 1)
}

func TestHistorySize(t *testing.T) {
	cmdA := new(command.MockCommand)
	cmdB := new(command.MockCommand)
	cmdA.On("Execute").Return(nil)
	cmdB.On("Execute").Return(nil)
	cmdB.On("Unexecute").Return(nil)

	history := New(sampleSize)
	history.Record(cmdA)
	history.Record(cmdB)
	history.Undo()
	assert.False(t, history.CanUndo())

	cmdA.AssertNumberOfCalls(t, "Execute", 1)
	cmdB.AssertNumberOfCalls(t, "Execute", 1)
	cmdB.AssertNumberOfCalls(t, "Unexecute", 1)
}
