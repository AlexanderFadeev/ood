package main

import (
	"fmt"

	"ood/cli_util"
	"ood/lab8/gumball_machine"
)

const startingGumballsCount = 5

func main() {
	if cli_util.PromtYesNo("Should we use multi gumball machine?") {
		playWithGumballMachine(gumball_machine.NewMulti(startingGumballsCount))
	} else {
		playWithGumballMachine(gumball_machine.New(startingGumballsCount))
	}
}

func playWithGumballMachine(m gumball_machine.GumballMachine) {
	fmt.Println(m)

	m.InsertQuarter()
	m.TurnCrank()
	fmt.Println(m)

	m.InsertQuarter()
	m.EjectQuarter()
	m.TurnCrank()
	fmt.Println(m)

	m.InsertQuarter()
	m.TurnCrank()
	m.InsertQuarter()
	m.TurnCrank()
	m.EjectQuarter()
	fmt.Println(m)

	m.InsertQuarter()
	m.InsertQuarter()
	m.TurnCrank()
	m.InsertQuarter()
	m.TurnCrank()
	m.InsertQuarter()
	m.TurnCrank()
	fmt.Println(m)
}
