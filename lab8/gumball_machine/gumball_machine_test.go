package gumball_machine

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

const startCount = 3
const refillCount = 2

type GumballMachineTestSuite struct {
	suite.Suite
	machine RefillableGumballMachine
	buf     *bytes.Buffer
}

func TestGumballMachineTestSuite(t *testing.T) {
	suite.Run(t, new(GumballMachineTestSuite))
}

func (s *GumballMachineTestSuite) SetupTest() {
	s.buf = &bytes.Buffer{}
	s.machine = newGumballMachine(startCount, s.buf)
}

func (s *GumballMachineTestSuite) AssertBufEqual(str string) {
	s.Equal(str, s.buf.String())
	s.buf.Reset()
}

func (s *GumballMachineTestSuite) AssertBufIsEmpty() {
	s.AssertBufEqual("")
}

func (s *GumballMachineTestSuite) Empty(count uint) {
	for i := uint(0); i < count-1; i++ {
		s.machine.InsertQuarter()
		s.AssertBufEqual("You inserted a quarter\n")
		s.machine.TurnCrank()
		s.AssertBufEqual("You turned a crank\nA gumball comes rolling out the slot...\n")
	}

	s.machine.InsertQuarter()
	s.AssertBufEqual("You inserted a quarter\n")
	s.machine.TurnCrank()
	s.AssertBufEqual("You turned a crank\nA gumball comes rolling out the slot...\nOops, out of gumballs\n")
	s.Equal("Gumball machine, 0 gumballs, sold out", s.machine.String())
}

func (s *GumballMachineTestSuite) TestNormalUseCase() {
	s.AssertBufIsEmpty()
	str := fmt.Sprintf("Gumball machine, %d gumballs, waiting for a quarter", startCount)
	s.Equal(str, s.machine.String())

	s.machine.InsertQuarter()
	s.AssertBufEqual("You inserted a quarter\n")
	str = fmt.Sprintf("Gumball machine, %d gumballs, waiting for turn of crank", startCount)
	s.Equal(str, s.machine.String())

	s.machine.TurnCrank()
	s.AssertBufEqual("You turned a crank\nA gumball comes rolling out the slot...\n")
	str = fmt.Sprintf("Gumball machine, %d gumballs, waiting for a quarter", startCount-1)
	s.Equal(str, s.machine.String())
}

func (s *GumballMachineTestSuite) TestInsertNoSpaceForQuarter() {
	s.machine.InsertQuarter()
	s.buf.Reset()
	s.machine.InsertQuarter()
	s.AssertBufEqual("You can't insert another quarter\n")
}

func (s *GumballMachineTestSuite) TestTurnNoQuarter() {
	s.machine.TurnCrank()
	s.AssertBufEqual("You turned but there's no quarter\n")
}

func (s *GumballMachineTestSuite) TestEjectQuarter() {
	s.machine.InsertQuarter()
	s.buf.Reset()
	s.machine.EjectQuarter()
	s.AssertBufEqual("Quarter returned\n")
}

func (s *GumballMachineTestSuite) TestEjectNoQuarter() {
	s.machine.EjectQuarter()
	s.AssertBufEqual("You haven't inserted a quarter\n")
}

func (s *GumballMachineTestSuite) TestRefill() {
	s.machine.Refill(refillCount)
	s.Empty(startCount + refillCount)
}

func (s *GumballMachineTestSuite) TestRefillWithQuarter() {
	s.machine.InsertQuarter()
	s.machine.Refill(refillCount)
	s.buf.Reset()
	s.machine.TurnCrank()
	s.AssertBufEqual("You turned a crank\nA gumball comes rolling out the slot...\n")

	s.Empty(startCount + refillCount - 1)
}

func (s *GumballMachineTestSuite) TestRefillEmpty() {
	s.Empty(startCount)
	s.machine.Refill(refillCount)
	s.Empty(refillCount)
}

func (s *GumballMachineTestSuite) TestSoldOut() {
	s.Empty(startCount)
	s.machine.EjectQuarter()
	s.AssertBufEqual("You can't eject, you haven't inserted a quarter yet\n")
	s.machine.InsertQuarter()
	s.AssertBufEqual("You can't insert a quarter, the machine is sold out\n")
	s.machine.TurnCrank()
	s.AssertBufEqual("You turned but there's no gumballs\n")
}
