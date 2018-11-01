package editor

import (
	"fmt"
	"io"
	"strings"

	"ood/lab5/document"

	"github.com/pkg/errors"
)

type Editor interface {
	InsertParagraph([]string, io.Writer) error
	InsertImage([]string, io.Writer) error
	SetTitle([]string, io.Writer) error
	List([]string, io.Writer) error
	ReplaceText([]string, io.Writer) error
	ResizeImage([]string, io.Writer) error
	DeleteItem([]string, io.Writer) error
	Undo() error
	Redo() error
	Save([]string, io.Writer) error
	Release() error
}

type editor struct {
	doc document.Document
}

func New() (Editor, error) {
	doc, err := document.New()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create a document")
	}

	return &editor{
		doc: doc,
	}, nil
}

func (e *editor) InsertParagraph(args []string, _ io.Writer) error {
	if len(args) < 2 {
		return errors.New("Not enough args")
	}

	text := strings.Join(args[1:], " ")

	var pos int
	_, err := fmt.Sscan(args[0], &pos)
	if err == nil {
		return e.doc.InsertParagraph(text, pos)
	}

	if args[0] != "end" {
		fmt.Println(args[0])
		return errors.New("Expected position to be a number or `end`")
	}

	return e.doc.AddParagraph(text)
}

func (e *editor) InsertImage(args []string, _ io.Writer) error {
	if len(args) < 4 {
		return errors.New("Not enough args")
	}

	var width, height, pos int
	_, err := fmt.Sscan(args[1], &width)
	if err != nil {
		return errors.Wrap(err, "Invalid width")
	}

	_, err = fmt.Sscan(args[2], &height)
	if err != nil {
		return errors.Wrap(err, "Invalid height")
	}

	_, err = fmt.Sscan(args[0], &pos)
	if err == nil {
		return e.doc.InsertImage(args[3], width, height, pos)
	}

	if args[0] != "end" {
		return errors.New("Expected position to be a number or `end`")
	}

	return e.doc.AddImage(args[3], width, height)
}

func (e *editor) SetTitle(args []string, _ io.Writer) error {
	title := strings.Join(args, " ")
	return e.doc.SetTitle(title)
}

func (e *editor) List(_ []string, w io.Writer) error {
	io.WriteString(w, fmt.Sprintf("%s\n", e.doc.String()))
	return nil
}

func (e *editor) ReplaceText(args []string, _ io.Writer) error {
	if len(args) < 2 {
		return errors.New("Not enough args")
	}

	var pos int
	_, err := fmt.Sscan(args[0], &pos)
	if err != nil {
		return errors.Wrap(err, "Invalid position")
	}

	text := strings.Join(args[1:], " ")
	return e.doc.EditParagraph(pos, text)
}

func (e *editor) ResizeImage(args []string, _ io.Writer) error {
	if len(args) < 3 {
		return errors.New("Not enough args")
	}

	var width, height, pos int
	_, err := fmt.Sscan(args[0], &pos)
	if err != nil {
		return errors.Wrap(err, "Invalid pos")
	}

	_, err = fmt.Sscan(args[1], &width)
	if err != nil {
		return errors.Wrap(err, "Invalid width")
	}

	_, err = fmt.Sscan(args[2], &height)
	if err != nil {
		return errors.Wrap(err, "Invalid height")
	}

	return e.doc.ResizeImage(pos, width, height)
}

func (e *editor) DeleteItem(args []string, _ io.Writer) error {
	if len(args) < 1 {
		return errors.New("Not enough args")
	}

	var pos int
	_, err := fmt.Sscan(args[0], &pos)
	if err != nil {
		return errors.Wrap(err, "Failed to parse position")
	}

	return e.doc.DeleteElement(pos)
}

func (e *editor) Save(args []string, _ io.Writer) error {
	if len(args) < 1 {
		return errors.New("Not enough args")
	}

	path := strings.Join(args, " ")
	return e.doc.Save(path)
}

func (e *editor) Redo() error {
	if !e.doc.CanRedo() {
		return errors.New("Can't redo")
	}

	return e.doc.Redo()
}

func (e *editor) Undo() error {
	if !e.doc.CanUndo() {
		return errors.New("Can't undo")
	}

	return e.doc.Undo()
}

func (e *editor) Release() error {
	return e.doc.Release()
}
