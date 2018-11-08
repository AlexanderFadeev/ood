package gumball_machine

import (
	"fmt"
	"io"
	"os"
)

type naive struct {
	ballsCount       uint
	quartersCount    uint
	quartersCapacity uint
	writer           io.Writer
}

func newNaive(count uint, quartersCapacity uint, writer io.Writer) RefillableGumballMachine {
	return &naive{
		ballsCount:       count,
		quartersCapacity: quartersCapacity,
		writer:           writer,
	}
}

func NewNaive(count uint) RefillableGumballMachine {
	return newGumballMachine(count, 1, os.Stdout)
}

func NewNaiveMulti(count uint) RefillableGumballMachine {
	return newGumballMachine(count, multiGumballMachineQuartersCapacity, os.Stdout)
}

func (n *naive) String() string {
	var stateStr string
	if n.ballsCount == 0 {
		stateStr = "sold out"
	} else if n.quartersCount > 0 {
		stateStr = "waiting for turn of crank"
	} else {
		stateStr = "waiting for a quarter"
	}

	return fmt.Sprintf("Gumball machine, %d gumballs, %s", n.ballsCount, stateStr)
}

func (n *naive) InsertQuarter() {
	if n.ballsCount == 0 {
		n.println("You can't insert a quarter, the machine is sold out")
		return
	}
	if n.quartersCount == n.quartersCapacity {
		n.println("You can't insert another quarter")
		return
	}
	n.quartersCount++
	n.println("You inserted a quarter")
}

func (n *naive) EjectQuarter() {
	if n.ballsCount == 0 && n.quartersCount == 0 {
		n.println("You can't eject, you haven't inserted a quarter yet")
		return
	}
	if n.quartersCount == 0 {
		n.println("You haven't inserted a quarter")
		return
	}
	n.quartersCount--
	n.println("Quarter returned")
}

func (n *naive) TurnCrank() {
	if n.ballsCount == 0 {
		n.println("You turned but there's no gumballs")
		return
	}
	if n.quartersCount == 0 {
		n.println("You turned but there's no quarter")
		return
	}
	n.println("You turned a crank")
	n.quartersCount--
	n.ballsCount--
	n.println("A gumball comes rolling out the slot...")
	if n.ballsCount == 0 {
		n.println("Oops, out of gumballs")
	}
}

func (n *naive) Refill(count uint) {
	n.ballsCount += count
}

func (n *naive) println(args ...interface{}) {
	fmt.Fprintln(n.writer, args...)
}
