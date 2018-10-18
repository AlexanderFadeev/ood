package menu

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

type Handler func([]string, io.Writer) error
type HandlerWithoutArgs func() error

type Menu interface {
	AddCommand(name, description string, handler Handler) error
	AddCommandWithoutArgs(name, description string, handler HandlerWithoutArgs) error
	SetDefaultHandler(HandlerWithoutArgs)
	Run()
	Help()
	Exit()
}

type menu struct {
	reader         io.Reader
	writer         io.Writer
	commands       map[string]commandInfo
	defaultHandler HandlerWithoutArgs
	shouldExit     bool
}

func New(reader io.Reader, writer io.Writer) Menu {
	return &menu{
		reader:         reader,
		writer:         writer,
		defaultHandler: func() error { return nil },
		commands:       make(map[string]commandInfo),
		shouldExit:     false,
	}
}

type commandInfo struct {
	description string
	handler     Handler
}

func (m *menu) AddCommand(name, description string, handler Handler) error {
	if _, ok := m.commands[name]; ok {
		return errors.Errorf("Command `%s` is already added to menu", name)
	}

	m.commands[name] = commandInfo{
		description: description,
		handler:     handler,
	}

	return nil
}

func (m *menu) AddCommandWithoutArgs(name, description string, handler HandlerWithoutArgs) error {
	return m.AddCommand(name, description, wrapHandlerWithoutArgs(handler))
}

func (m *menu) SetDefaultHandler(defaultHandler HandlerWithoutArgs) {
	m.defaultHandler = defaultHandler
}

func (m *menu) Run() {
	scanner := bufio.NewScanner(m.reader)
	for !m.shouldExit && m.promt() && scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, " ")
		if len(args) == 0 {
			continue
		}

		err := m.runCommand(args)
		if err != nil {
			io.WriteString(m.writer, fmt.Sprintf("%s\n", err.Error()))
		}
	}
}

func (m *menu) Help() {
	io.WriteString(m.writer, "Commands list:\n")
	for name, info := range m.commands {
		desc := fmt.Sprintf("\t%s: %s\n", name, info.description)
		io.WriteString(m.writer, desc)
	}
}

func (m *menu) Exit() {
	m.shouldExit = true
}

func (m *menu) runCommand(args []string) error {
	command, ok := m.commands[args[0]]
	if !ok {
		return m.defaultHandler()
	}

	return command.handler(args[1:], m.writer)
}

func (m *menu) promt() bool {
	io.WriteString(m.writer, ">")
	return true
}

func wrapHandlerWithoutArgs(f HandlerWithoutArgs) Handler {
	return func(_ []string, _ io.Writer) error {
		return f()
	}
}
