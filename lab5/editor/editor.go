package editor

import (
	"io/ioutil"
	"path"

	"ood/lab5/command"
	"ood/lab5/document"
	"ood/lab5/history"
	"ood/lab5/storage"

	"github.com/pkg/errors"
)

const maxHistoryLength = 10

type Editor interface {
	SetTitle(title string) error
	List() string

	InsertParagraph(text string, pos int) error
	AddParagraph(text string) error
	EditParagraph(pos int, text string) error

	InsertImage(path string, width, height int, pos int) error
	AddImage(path string, width, height int) error
	ResizeImage(pos int, width, height int) error

	DeleteElement(pos int) error

	Save(path string) error

	CanRedo() bool
	CanUndo() bool
	Redo() error
	Undo() error
}

type editor struct {
	doc            document.Document
	history        history.History
	workDirStorage storage.Storage
	globalStorage  storage.Storage
	tempStorage    storage.Storage
}

func New() (Editor, error) {
	wdStorage, err := storage.NewLocal(".")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create workdir storage")
	}

	tempStorage, err := storage.NewTemp()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create temp storage")
	}

	globalStorage, err := storage.NewLocal("")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create global storage")
	}

	return &editor{
		doc:            document.New(),
		history:        history.New(maxHistoryLength),
		workDirStorage: wdStorage,
		globalStorage:  globalStorage,
		tempStorage:    tempStorage,
	}, nil
}

func (e *editor) SetTitle(title string) error {
	oldTitle := e.doc.GetTitle()

	cmd := command.New(func() error {
		e.doc.SetTitle(title)
		return nil
	}, func() error {
		e.doc.SetTitle(oldTitle)
		return nil
	})

	err := e.history.AddAndExecute(cmd)
	return errors.Wrap(err, "Failed to add and execute command")
}

func (e *editor) List() string {
	return e.doc.String()
}

func (e *editor) InsertParagraph(text string, pos int) error {
	paragraph := document.NewParagraph(text)

	cmd := command.New(func() error {
		return e.doc.InsertElement(paragraph, pos)
	}, func() error {
		return e.doc.DeleteElement(pos)
	})

	err := e.history.AddAndExecute(cmd)
	return errors.Wrap(err, "Failed to add and execute command")
}

func (e *editor) AddParagraph(text string) error {
	return e.InsertParagraph(text, e.doc.GetElementsCount())
}

func (e *editor) EditParagraph(pos int, text string) error {
	elem, err := e.doc.GetElement(pos)
	if err != nil {
		return errors.Wrap(err, "Failed to get element")
	}

	paragraph, ok := elem.(document.Paragraph)
	if !ok {
		return errors.Errorf("Element at position %d is not a paragraph", pos)
	}

	oldText := paragraph.GetText()

	cmd := command.New(func() error {
		paragraph.SetText(text)
		return nil
	}, func() error {
		paragraph.SetText(oldText)
		return nil
	})

	err = e.history.AddAndExecute(cmd)
	return errors.Wrap(err, "Failed to add and execute command")
}

func (e *editor) InsertImage(path string, width, height int, pos int) error {
	cmd, err := e.newInsertImageCommand(pos, width, height, path)
	if err != nil {
		return errors.Wrap(err, "Failed to create command")
	}

	err = e.history.AddAndExecute(cmd)
	return errors.Wrap(err, "Failed to add and execute command")
}

func (e *editor) AddImage(path string, width, height int) error {
	return e.InsertImage(path, width, height, e.doc.GetElementsCount())
}

func (e *editor) ResizeImage(pos int, width, height int) error {
	elem, err := e.doc.GetElement(pos)
	if err != nil {
		return errors.Wrap(err, "Failed to get element")
	}

	img, ok := elem.(document.Image)
	if !ok {
		return errors.Errorf("Element at position %d is not an image", pos)
	}

	oldWidth, oldHeight := img.GetSize()

	cmd := command.New(func() error {
		img.SetSize(width, height)
		return nil
	}, func() error {
		img.SetSize(oldWidth, oldHeight)
		return nil
	})

	err = e.history.AddAndExecute(cmd)
	return errors.Wrap(err, "Failed to add and execute command")
}

func (e *editor) DeleteElement(pos int) error {
	elem, err := e.doc.GetElement(pos)
	if err != nil {
		return errors.Wrap(err, "Failed to get element")
	}

	cmd := command.New(func() error {
		return e.doc.DeleteElement(pos)
	}, func() error {
		return e.doc.InsertElement(elem, pos)
	})

	err = e.history.AddAndExecute(cmd)
	return errors.Wrap(err, "Failed to add and execute command")
}

func (e *editor) Save(htmlPath string) error {
	dir := path.Dir(htmlPath)
	htmlStorage, err := storage.NewLocal(dir)
	if err != nil {
		return errors.Wrap(err, "Failed to create HTML storage")
	}

	err = storage.CopyAll(e.tempStorage, htmlStorage)
	if err != nil {
		return errors.Wrap(err, "Failed to copy files to HTML storage")
	}

	err = ioutil.WriteFile(htmlPath, []byte(e.doc.ToHTML()), 0x666)
	return errors.Wrap(err, "Failed to write html file")
}

func (e *editor) CanRedo() bool {
	return e.history.CanRedo()
}

func (e *editor) CanUndo() bool {
	return e.history.CanUndo()
}

func (e *editor) Redo() error {
	if !e.history.CanRedo() {
		return errors.New("Cannot redo")
	}

	err := e.history.Redo()
	return errors.Wrap(err, "Failed to redo")
}

func (e *editor) Undo() error {
	if !e.history.CanUndo() {
		return errors.New("Cannot undo")
	}

	err := e.history.Undo()
	return errors.Wrap(err, "Failed to undo")
}
