package gumball_machine

import "io"

type machineFactory interface {
	createMachine(count uint, quartersCapacity uint, writer io.Writer) RefillableGumballMachine
}

type commonMachineFactory struct{}

func (commonMachineFactory) createMachine(count uint, quartersCapacity uint, writer io.Writer) RefillableGumballMachine {
	return newGumballMachine(count, quartersCapacity, writer)
}

type naiveMachineFactory struct{}

func (naiveMachineFactory) createMachine(count uint, quartersCapacity uint, writer io.Writer) RefillableGumballMachine {
	return newNaive(count, quartersCapacity, writer)
}
