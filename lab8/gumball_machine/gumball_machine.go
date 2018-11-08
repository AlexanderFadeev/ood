package gumball_machine

import (
	"fmt"
	"io"
	"os"
)

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
	state      state
	ballsCount uint
	writer     io.Writer
}

func newGumballMachine(count uint, writer io.Writer) RefillableGumballMachine {
	var state state = new(stateNoQuarter)
	if count == 0 {
		state = new(stateSoldOut)
	}

	return &gumballMachine{
		ballsCount: count,
		state:      state,
		writer:     writer,
	}
}

func New(count uint) RefillableGumballMachine {
	return newGumballMachine(count, os.Stdout)
}

func (n *gumballMachine) String() string {
	return fmt.Sprintf("Gumball machine, %d gumballs, %s", n.ballsCount, n.state)
}

func (n *gumballMachine) InsertQuarter() {
	n.state.insertQuarter(n)
}

func (n *gumballMachine) EjectQuarter() {
	n.state.ejectQuarter(n)
}

func (n *gumballMachine) TurnCrank() {
	n.state.turnCrank(n)
}

func (n *gumballMachine) Refill(count uint) {
	n.state.refill(n, count)
}

func (n *gumballMachine) setState(state state) {
	n.state = state
}

func (n *gumballMachine) incBallsCount(count uint) {
	n.ballsCount += count
}

func (n *gumballMachine) getBallsCount() uint {
	return n.ballsCount
}

func (n *gumballMachine) releaseBall() {
	n.ballsCount--
	n.println("A gumball comes rolling out the slot...")
}

func (n *gumballMachine) println(args ...interface{}) {
	fmt.Fprintln(n.writer, args...)
}
