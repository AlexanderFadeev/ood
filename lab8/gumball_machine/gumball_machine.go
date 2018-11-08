package gumball_machine

import (
	"fmt"
	"io"
	"os"
)

const multiGumballMachineQuartersCapacity = 5

type GumballMachine interface {
	fmt.Stringer

	InsertQuarter()
	EjectQuarter()
	TurnCrank()
}

type RefillableGumballMachine interface {
	GumballMachine

	Refill(uint)
}
type gumballMachine struct {
	state            state
	ballsCount       uint
	quartersCount    uint
	quartersCapacity uint
	writer           io.Writer
}

func newGumballMachine(count uint, quartersCapacity uint, writer io.Writer) RefillableGumballMachine {
	var state state = new(stateNoQuarter)
	if count == 0 {
		state = new(stateSoldOut)
	}

	return &gumballMachine{
		ballsCount:       count,
		state:            state,
		quartersCapacity: quartersCapacity,
		writer:           writer,
	}
}

func New(count uint) RefillableGumballMachine {
	return newGumballMachine(count, 1, os.Stdout)
}

func NewMulti(count uint) RefillableGumballMachine {
	return newGumballMachine(count, multiGumballMachineQuartersCapacity, os.Stdout)
}

func (gm *gumballMachine) String() string {
	return fmt.Sprintf("Gumball machine, %d gumballs, %s", gm.ballsCount, gm.state)
}

func (gm *gumballMachine) InsertQuarter() {
	gm.state.insertQuarter(gm)
}

func (gm *gumballMachine) EjectQuarter() {
	gm.state.ejectQuarter(gm)
}

func (gm *gumballMachine) TurnCrank() {
	gm.state.turnCrank(gm)
}

func (gm *gumballMachine) Refill(count uint) {
	gm.state.refill(gm, count)
}

func (gm *gumballMachine) setState(state state) {
	gm.state = state
}

func (gm *gumballMachine) incBallsCount(count uint) {
	gm.ballsCount += count
}

func (gm *gumballMachine) getBallsCount() uint {
	return gm.ballsCount
}

func (gm *gumballMachine) getQuartersCapacity() uint {
	return gm.quartersCapacity
}

func (gm *gumballMachine) getQuartersCount() uint {
	return gm.quartersCount
}

func (gm *gumballMachine) changeQuartersCount(count int) {
	gm.quartersCount = uint(int(gm.quartersCount) + count)
}

func (gm *gumballMachine) releaseBall() {
	gm.ballsCount--
	gm.println("A gumball comes rolling out the slot...")
}

func (gm *gumballMachine) println(args ...interface{}) {
	fmt.Fprintln(gm.writer, args...)
}
