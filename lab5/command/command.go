package command

type Command interface {
	Execute() error
	Unexecute() error
	Release() error
}
