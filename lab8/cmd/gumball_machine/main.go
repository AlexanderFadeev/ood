package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/AlexanderFadeev/ood/cli_util"
	"github.com/AlexanderFadeev/ood/lab5/menu"
	"github.com/AlexanderFadeev/ood/lab8/gumball_machine"

	"github.com/pkg/errors"
)

const startingGumballsCount = 5

func main() {
	if cli_util.PromtYesNo("Should we use multi gumball machine?") {
		useGumballMachine(gumball_machine.NewMulti(startingGumballsCount))
	} else {
		useGumballMachine(gumball_machine.New(startingGumballsCount))
	}
}

func useGumballMachine(m gumball_machine.RefillableGumballMachine) {
	if cli_util.PromtYesNo("Should we use menu to interact with a machine?") {
		interactGumballMachineViaMenu(m)
	} else {
		playWithGumballMachine(m)
	}
}

func interactGumballMachineViaMenu(machine gumball_machine.RefillableGumballMachine) {
	m := menu.New(os.Stdin, os.Stdout)
	m.AddCommandWithoutArgs("InsertQuarter", "Insert a quarter", wrapSimpleFunc(machine.InsertQuarter))
	m.AddCommandWithoutArgs("EjectQuarter", "Eject a quarter", wrapSimpleFunc(machine.EjectQuarter))
	m.AddCommandWithoutArgs("TurnCrank", "Turn the crank", wrapSimpleFunc(machine.TurnCrank))
	m.AddCommand("Refill", "Refill the gumball machine", makeRefillFunc(machine))
	m.AddCommandWithoutArgs("Status", "Show gumball machine status", func() error {
		fmt.Println(machine)
		return nil
	})
	m.AddCommandWithoutArgs("Help", "Show help menu", wrapSimpleFunc(m.Help))
	m.AddCommandWithoutArgs("Exit", "Exit", wrapSimpleFunc(m.Exit))
	m.SetDefaultHandler(func() error {
		fmt.Println("Use `Help` command to show help menu")
		return nil
	})

	m.Run()
}

func wrapSimpleFunc(fn func()) menu.HandlerWithoutArgs {
	return func() error {
		fn()
		return nil
	}
}

func makeRefillFunc(machine gumball_machine.RefillableGumballMachine) menu.Handler {
	return func(args []string, writer io.Writer) error {
		if len(args) == 0 {
			return errors.New("Invalid args count")
		}

		count, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return errors.Wrap(err, "Failed to parse args")
		}
		if count < 0 {
			return errors.New("Expected non-negative value")
		}

		machine.Refill(uint(count))
		return nil
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
