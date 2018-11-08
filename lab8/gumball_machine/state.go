package gumball_machine

import (
	"fmt"
)

type state interface {
	fmt.Stringer

	insertQuarter(context)
	ejectQuarter(context)
	turnCrank(context)
	refill(context, uint)
}

type context interface {
	GumballMachine

	setState(state)
	incBallsCount(uint)
	getBallsCount() uint
	releaseBall()
	println(...interface{})
}

type stateSoldOut struct{}
type stateNoQuarter struct{}
type stateHasQuarter struct{}

//Sold out state

func (stateSoldOut) String() string {
	return "sold out"
}

func (stateSoldOut) insertQuarter(ctx context) {
	ctx.println("You can't insert a quarter, the machine is sold out")
}

func (stateSoldOut) ejectQuarter(ctx context) {
	ctx.println("You can't eject, you haven't inserted a quarter yet")
}

func (stateSoldOut) turnCrank(ctx context) {
	ctx.println("You turned but there's no gumballs")
}

func (stateSoldOut) refill(ctx context, count uint) {
	ctx.incBallsCount(count)
	if count > 0 {
		ctx.setState(new(stateNoQuarter))
	}
}

//No quarter state

func (stateNoQuarter) String() string {
	return "waiting for a quarter"
}

func (stateNoQuarter) insertQuarter(ctx context) {
	ctx.println("You inserted a quarter")
	ctx.setState(new(stateHasQuarter))
}

func (stateNoQuarter) ejectQuarter(ctx context) {
	ctx.println("You haven't inserted a quarter")
}

func (stateNoQuarter) turnCrank(ctx context) {
	ctx.println("You turned but there's no quarter")
}

func (stateNoQuarter) refill(ctx context, count uint) {
	ctx.incBallsCount(count)
}

//Has quarter state

func (stateHasQuarter) String() string {
	return "waiting for turn of crank"
}

func (stateHasQuarter) insertQuarter(ctx context) {
	ctx.println("You can't insert another quarter")
}

func (stateHasQuarter) ejectQuarter(ctx context) {
	ctx.println("Quarter returned")
	ctx.setState(new(stateNoQuarter))
}

func (stateHasQuarter) turnCrank(ctx context) {
	ctx.println("You turned a crank")
	ctx.releaseBall()
	if ctx.getBallsCount() == 0 {
		ctx.println("Oops, out of gumballs")
		ctx.setState(new(stateSoldOut))
	} else {
		ctx.setState(new(stateNoQuarter))
	}
}

func (stateHasQuarter) refill(ctx context, count uint) {
	ctx.incBallsCount(count)
}
