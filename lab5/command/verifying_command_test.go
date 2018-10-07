package command

//go:generate mockery -dir ./ -name=Command -inpkg -case=underscore

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var sampleError = errors.New("Sample error")

func newMock() *MockCommand {
	mock := new(MockCommand)
	mock.On("Execute").Return(nil)
	mock.On("Unexecute").Return(nil)
	mock.On("Release").Return(nil)
	return mock
}

func assertNumberOfCalls(t *testing.T, mock *MockCommand, executed, unexecuted, released int) {
	mock.AssertNumberOfCalls(t, "Execute", executed)
	mock.AssertNumberOfCalls(t, "Unexecute", unexecuted)
	mock.AssertNumberOfCalls(t, "Release", released)
}

func TestVerifyingCommandExecuteTwice(t *testing.T) {
	mock := newMock()
	wrapper := WrapVerify(mock)

	wrapper.Execute()
	assert.Panics(t, func() { wrapper.Execute() })

	assertNumberOfCalls(t, mock, 1, 0, 0)
}

func TestVerifyingCommandExecuteAndRelease(t *testing.T) {
	mock := newMock()
	wrapper := WrapVerify(mock)

	wrapper.Execute()
	assert.Panics(t, func() { wrapper.Release() })

	assertNumberOfCalls(t, mock, 1, 0, 0)
}

func TestVerifyingCommandUnexecuteTwice(t *testing.T) {
	mock := newMock()
	wrapper := WrapVerify(mock)
	wrapper.Execute()

	wrapper.Unexecute()
	assert.Panics(t, func() { wrapper.Unexecute() })

	assertNumberOfCalls(t, mock, 1, 1, 0)
}

func TestVerifyingCommandReleaseTwice(t *testing.T) {
	mock := newMock()
	wrapper := WrapVerify(mock)

	wrapper.Release()
	assert.Panics(t, func() { wrapper.Release() })

	assertNumberOfCalls(t, mock, 0, 0, 1)
}

func TestVerifyingCommandReleaseAndExecute(t *testing.T) {
	mock := newMock()
	wrapper := WrapVerify(mock)

	wrapper.Release()
	assert.Panics(t, func() { wrapper.Execute() })

	assertNumberOfCalls(t, mock, 0, 0, 1)
}

func TestVerifyingCommandReleaseAndUnexecute(t *testing.T) {
	mock := newMock()
	wrapper := WrapVerify(mock)

	wrapper.Release()
	assert.Panics(t, func() { wrapper.Unexecute() })

	assertNumberOfCalls(t, mock, 0, 0, 1)
}

func TestVerifyingCommandExecuteError(t *testing.T) {
	mock := new(MockCommand)
	mock.On("Execute").Return(sampleError)
	wrapper := WrapVerify(mock)

	err := wrapper.Execute()
	assert.NotNil(t, err)
}

func TestVerifyingCommandUnexecuteError(t *testing.T) {
	mock := new(MockCommand)
	mock.On("Execute").Return(nil)
	mock.On("Unexecute").Return(sampleError)
	wrapper := WrapVerify(mock)
	wrapper.Execute()

	err := wrapper.Unexecute()
	assert.NotNil(t, err)
}

func TestVerifyingCommandReleaseError(t *testing.T) {
	mock := new(MockCommand)
	mock.On("Release").Return(sampleError)
	wrapper := WrapVerify(mock)

	err := wrapper.Release()
	assert.NotNil(t, err)
}
