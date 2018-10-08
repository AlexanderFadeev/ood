package history

import (
	"testing"

	"ood/lab5/command"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const sampleSize = 1

var sampleError = errors.New("Sample error")

func TestAddAndExecute(t *testing.T) {
	com := new(command.MockCommand)
	com.On("Execute").Return(nil)

	history := New(sampleSize)
	err := history.AddAndExecute(com)
	assert.Nil(t, err)

	com.AssertNumberOfCalls(t, "Execute", 1)
}

func TestAddAndExecuteError(t *testing.T) {
	com := new(command.MockCommand)
	com.On("Execute").Return(sampleError)

	history := New(sampleSize)
	err := history.AddAndExecute(com)
	assert.NotNil(t, err)

	com.AssertNumberOfCalls(t, "Execute", 1)
}

func TestUndo(t *testing.T) {
	com := new(command.MockCommand)
	com.On("Execute").Return(nil)
	com.On("Unexecute").Return(nil)

	history := New(sampleSize)
	assert.False(t, history.CanUndo())
	assert.Panics(t, func() { history.Undo() })

	history.AddAndExecute(com)
	assert.True(t, history.CanUndo())

	history.Undo()
	assert.False(t, history.CanUndo())
	assert.Panics(t, func() { history.Undo() })

	com.AssertNumberOfCalls(t, "Execute", 1)
	com.AssertNumberOfCalls(t, "Unexecute", 1)
}

func TestUndoError(t *testing.T) {
	com := new(command.MockCommand)
	com.On("Execute").Return(nil)
	com.On("Unexecute").Return(sampleError)

	history := New(sampleSize)
	history.AddAndExecute(com)
	err := history.Undo()
	assert.NotNil(t, err)

	com.AssertNumberOfCalls(t, "Execute", 1)
	com.AssertNumberOfCalls(t, "Unexecute", 1)
}

func TestRedo(t *testing.T) {
	com := new(command.MockCommand)
	com.On("Execute").Return(nil)
	com.On("Unexecute").Return(nil)

	history := New(sampleSize)
	assert.False(t, history.CanRedo())
	assert.Panics(t, func() { history.Redo() })

	history.AddAndExecute(com)
	assert.False(t, history.CanRedo())
	assert.Panics(t, func() { history.Redo() })

	history.Undo()
	assert.True(t, history.CanRedo())

	history.Redo()
	assert.False(t, history.CanRedo())
	assert.Panics(t, func() { history.Redo() })

	com.AssertNumberOfCalls(t, "Execute", 2)
	com.AssertNumberOfCalls(t, "Unexecute", 1)
}

func TestRedoError(t *testing.T) {
	com := new(command.MockCommand)
	calledTime := 0

	com.On("Execute").Return(func() error {
		calledTime++
		if calledTime == 2 {
			return sampleError
		}
		return nil
	})
	com.On("Unexecute").Return(nil)

	history := New(sampleSize)
	history.AddAndExecute(com)
	history.Undo()
	err := history.Redo()
	assert.NotNil(t, err)

	com.AssertNumberOfCalls(t, "Execute", 2)
	com.AssertNumberOfCalls(t, "Unexecute", 1)
}

func TestAddAndExecuteAfterUndo(t *testing.T) {
	comA := new(command.MockCommand)
	comA.On("Execute").Return(nil)
	comA.On("Unexecute").Return(nil)
	comA.On("Release").Return(nil)
	comB := new(command.MockCommand)
	comB.On("Execute").Return(nil)

	history := New(sampleSize)
	history.AddAndExecute(comA)
	history.Undo()
	history.AddAndExecute(comB)

	comA.AssertNumberOfCalls(t, "Execute", 1)
	comA.AssertNumberOfCalls(t, "Unexecute", 1)
	comA.AssertNumberOfCalls(t, "Release", 1)
	comB.AssertNumberOfCalls(t, "Execute", 1)
}

func TestReleaseError(t *testing.T) {
	comA := new(command.MockCommand)
	comA.On("Execute").Return(nil)
	comA.On("Unexecute").Return(nil)
	comA.On("Release").Return(sampleError)
	comB := new(command.MockCommand)
	comB.On("Execute").Return(nil)

	history := New(sampleSize)
	history.AddAndExecute(comA)
	history.Undo()
	err := history.AddAndExecute(comB)
	assert.NotNil(t, err)

	comA.AssertNumberOfCalls(t, "Execute", 1)
	comA.AssertNumberOfCalls(t, "Unexecute", 1)
	comA.AssertNumberOfCalls(t, "Release", 1)
	comB.AssertNumberOfCalls(t, "Execute", 1)
}

func TestHistorySize(t *testing.T) {
	comA := new(command.MockCommand)
	comB := new(command.MockCommand)
	comA.On("Execute").Return(nil)
	comB.On("Execute").Return(nil)
	comB.On("Unexecute").Return(nil)

	history := New(sampleSize)
	history.AddAndExecute(comA)
	history.AddAndExecute(comB)
	history.Undo()
	assert.False(t, history.CanUndo())

	comA.AssertNumberOfCalls(t, "Execute", 1)
	comB.AssertNumberOfCalls(t, "Execute", 1)
	comB.AssertNumberOfCalls(t, "Unexecute", 1)
}
