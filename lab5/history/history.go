package history

import (
	"container/list"

	"github.com/AlexanderFadeev/ood/lab5/command"

	"github.com/pkg/errors"
)

type Recorder interface {
	Record(command.Command) error
}

type History interface {
	Recorder

	CanUndo() bool
	CanRedo() bool
	Undo() error
	Redo() error

	Release() error
}

type history struct {
	size               int
	list               *list.List
	lastCommandElement *list.Element
}

func New(size int) History {
	return &history{
		list: list.New(),
		size: size,
	}
}

func (h *history) Record(command command.Command) error {
	err := command.Execute()
	if err != nil {
		return errors.Wrap(err, "Failed to execute command")
	}

	eraseErr := h.eraseTail()

	h.list.PushBack(command)
	if h.list.Len() > h.size {
		h.list.Remove(h.list.Front())
	}
	h.lastCommandElement = h.list.Back()

	return errors.Wrap(eraseErr, "Failed to erase list tail")
}

func (h *history) CanUndo() bool {
	return h.lastCommandElement != nil
}

func (h *history) CanRedo() bool {
	return h.getNextCommandElement() != nil
}

func (h *history) Undo() error {
	err := h.getLastCommand().Unexecute()
	if err != nil {
		return errors.Wrap(err, "Failed to unexecute command")
	}

	h.lastCommandElement = h.lastCommandElement.Prev()
	return nil
}

func (h *history) Redo() error {
	err := h.getNextCommand().Execute()
	if err != nil {
		return errors.Wrap(err, "Failed to execute command")
	}

	h.lastCommandElement = h.getNextCommandElement()
	return nil
}

func (h *history) Release() error {
	for h.list.Back() != nil {
		cmd := h.list.Remove(h.list.Back()).(command.Command)
		err := cmd.Release()
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *history) eraseTail() (err error) {
	for h.list.Back() != h.lastCommandElement {
		cmd := h.list.Remove(h.list.Back()).(command.Command)
		currErr := cmd.Release()
		if currErr != nil {
			err = errors.Wrap(currErr, "Failed to release command")
		}
	}
	return err
}

func (h *history) getLastCommand() command.Command {
	return h.lastCommandElement.Value.(command.Command)
}

func (h *history) getNextCommandElement() *list.Element {
	if h.lastCommandElement == nil {
		return h.list.Front()
	}

	return h.lastCommandElement.Next()
}

func (h *history) getNextCommand() command.Command {
	return h.getNextCommandElement().Value.(command.Command)
}
