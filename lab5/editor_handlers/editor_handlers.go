package editor_handlers

import (
	"fmt"
	"io"
	"strings"

	"ood/lab5/editor"

	"github.com/pkg/errors"
)

type EditorHandlers interface {
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
}

type handlers struct {
	editor editor.Editor
}

func New(editor editor.Editor) EditorHandlers {
	return &handlers{
		editor: editor,
	}
}

func (h *handlers) InsertParagraph(args []string, _ io.Writer) error {
	if len(args) < 2 {
		return errors.New("Not enough args")
	}

	text := strings.Join(args[1:], " ")

	var pos int
	_, err := fmt.Sscan(args[0], &pos)
	if err == nil {
		return h.editor.InsertParagraph(text, pos)
	}

	if args[0] != "end" {
		fmt.Println(args[0])
		return errors.New("Expected position to be a number or `end`")
	}

	return h.editor.AddParagraph(text)
}

func (h *handlers) InsertImage(args []string, _ io.Writer) error {
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
		return h.editor.InsertImage(args[3], width, height, pos)
	}

	if args[0] != "end" {
		return errors.New("Expected position to be a number or `end`")
	}

	return h.editor.AddImage(args[3], width, height)
}

func (h *handlers) SetTitle(args []string, _ io.Writer) error {
	title := strings.Join(args, " ")
	return h.editor.SetTitle(title)
}

func (h *handlers) List(_ []string, w io.Writer) error {
	io.WriteString(w, fmt.Sprintf("%s\n", h.editor.List()))
	return nil
}

func (h *handlers) ReplaceText(args []string, _ io.Writer) error {
	if len(args) < 2 {
		return errors.New("Not enough args")
	}

	var pos int
	_, err := fmt.Sscan(args[0], &pos)
	if err != nil {
		return errors.Wrap(err, "Invalid position")
	}

	text := strings.Join(args[1:], " ")
	return h.editor.EditParagraph(pos, text)
}

func (h *handlers) ResizeImage(args []string, _ io.Writer) error {
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

	return h.editor.ResizeImage(pos, width, height)
}

func (h *handlers) DeleteItem(args []string, _ io.Writer) error {
	if len(args) < 1 {
		return errors.New("Not enough args")
	}

	var pos int
	_, err := fmt.Sscan(args[0], &pos)
	if err != nil {
		return errors.Wrap(err, "Failed to parse position")
	}

	return h.editor.DeleteElement(pos)
}

func (h *handlers) Save(args []string, _ io.Writer) error {
	if len(args) < 1 {
		return errors.New("Not enough args")
	}

	path := strings.Join(args, " ")
	return h.editor.Save(path)
}

func (h *handlers) Redo() error {
	if !h.editor.CanRedo() {
		return errors.New("Can't redo")
	}

	return h.editor.Redo()
}

func (h *handlers) Undo() error {
	if !h.editor.CanUndo() {
		return errors.New("Can't undo")
	}

	return h.editor.Undo()
}
