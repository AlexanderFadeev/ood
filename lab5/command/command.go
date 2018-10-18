package command

type Command interface {
	Execute() error
	Unexecute() error
	Release() error
}

type Callback func() error

type command struct {
	do, undo, release Callback
}

func noop() error {
	return nil
}

func New(do, undo Callback) Command {
	return WrapVerify(&command{do, undo, noop})
}

func NewWithRelease(do, undo, release Callback) Command {
	return WrapVerify(&command{do, undo, release})
}

func (c *command) Execute() error {
	return c.do()
}

func (c *command) Unexecute() error {
	return c.undo()
}

func (c *command) Release() error {
	return c.release()
}
