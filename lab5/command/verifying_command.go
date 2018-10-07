package command

type commandState int

const (
	stateUnexecuted commandState = iota
	stateExecuted
	stateReleased
)

type verifyingCommand struct {
	impl  Command
	state commandState
}

func WrapVerify(command Command) Command {
	return &verifyingCommand{
		impl:  command,
		state: stateUnexecuted,
	}
}

func (vc *verifyingCommand) Execute() error {
	if vc.state == stateExecuted {
		panic("Trying to execute already executed command")
	}
	if vc.state == stateReleased {
		panic("Trying to execute released command")
	}

	err := vc.impl.Execute()
	if err != nil {
		return err
	}

	vc.state = stateExecuted
	return nil
}

func (vc *verifyingCommand) Unexecute() error {
	if vc.state == stateUnexecuted {
		panic("Trying to unexecute already unexecuted command")
	}
	if vc.state == stateReleased {
		panic("Trying to unexecute released command")
	}

	err := vc.impl.Unexecute()
	if err != nil {
		return err
	}

	vc.state = stateUnexecuted
	return nil
}

func (vc *verifyingCommand) Release() error {
	if vc.state == stateReleased {
		panic("Trying to release already released command")
	}
	if vc.state == stateExecuted {
		panic("Trying to release executed command")
	}

	err := vc.impl.Release()
	if err != nil {
		return err
	}

	vc.state = stateReleased
	return nil
}
