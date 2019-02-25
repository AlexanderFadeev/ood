package document

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/AlexanderFadeev/ood/lab5/command"
	"github.com/AlexanderFadeev/ood/lab5/history"
	"github.com/AlexanderFadeev/ood/lab5/storage"
	"github.com/AlexanderFadeev/ood/lab5/vector"

	"github.com/pkg/errors"
)

const maxHistoryLength = 10

type Document interface {
	fmt.Stringer
	htmlFormatAcceptor

	GetTitle() string
	SetTitle(string) error

	InsertParagraph(text string, position int) error
	AddParagraph(text string) error
	EditParagraph(position int, text string) error

	InsertImage(path string, width, height int, pos int) error
	AddImage(path string, width, height int) error
	ResizeImage(pos int, width, height int) error

	GetElementsCount() int
	getElements() vector.Vector
	DeleteElement(position int) error
	Save(path string) error

	CanUndo() bool
	CanRedo() bool
	Undo() error
	Redo() error

	Release() error
}

type document struct {
	title    string
	elements vector.Vector

	history      history.History
	localStorage storage.Storage
	tempStorage  storage.TempStorage
}

func New() (Document, error) {
	tempStorage, err := storage.NewTemp()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create temp storage")
	}

	return &document{
		history:      history.New(maxHistoryLength),
		localStorage: storage.NewLocal(""),
		tempStorage:  tempStorage,
	}, nil
}

func (d *document) GetElementsCount() int {
	return len(d.elements)
}

func (d *document) GetTitle() string {
	return d.title
}

func (d *document) SetTitle(title string) error {
	oldTitle := d.title

	cmd := command.New(func() error {
		d.setTitle(title)
		return nil
	}, func() error {
		d.setTitle(oldTitle)
		return nil
	})

	err := d.history.Record(cmd)
	return errors.Wrap(err, "Failed to record the command")
}

func (d *document) setTitle(title string) {
	d.title = title
}

func (d *document) AddParagraph(text string) error {
	return d.InsertParagraph(text, d.GetElementsCount())
}

func (d *document) InsertParagraph(text string, position int) error {
	paragraph := newParagraph(text, d.history)
	return d.insertElement(paragraph, position)
}

func (d *document) EditParagraph(position int, text string) error {
	elem, err := d.GetElement(position)
	if err != nil {
		return errors.Wrap(err, "Failed to get element")
	}

	paragraph, ok := elem.(paragraph)
	if !ok {
		return errors.Errorf("element at position %d is not a paragraph", position)
	}

	return paragraph.setText(text)
}

func (d *document) InsertImage(path string, width, height int, pos int) error {
	img, err := d.newImage(path, width, height)
	if err != nil {
		return errors.Wrap(err, "Failed to create image")
	}

	cmd := command.NewWithRelease(func() error {
		return d.insertElementImpl(img, pos)
	}, func() error {
		return d.deleteElement(pos)
	}, func() error {
		return d.tempStorage.DeleteFile(img.getPath())
	})

	err = d.history.Record(cmd)
	return errors.Wrap(err, "Failed to record the command")
}

func (d *document) AddImage(path string, width, height int) error {
	return d.InsertImage(path, width, height, d.GetElementsCount())
}

func (d *document) ResizeImage(position int, width, height int) error {
	elem, err := d.GetElement(position)
	if err != nil {
		return errors.Wrap(err, "Failed to get element")
	}

	img, ok := elem.(image)
	if !ok {
		return errors.Errorf("element at position %d is not an image", position)
	}

	return img.setSize(width, height)
}

func (d *document) insertElement(element element, position int) error {
	cmd := command.New(func() error {
		return d.insertElementImpl(element, position)
	}, func() error {
		return d.deleteElement(position)
	})

	err := d.history.Record(cmd)
	return errors.Wrap(err, "Failed to record the command")
}

func (d *document) insertElementImpl(element element, position int) error {
	if !d.positionIsValid(position) && position != d.GetElementsCount() {
		return errors.Errorf("Invalid element position `%d`", position)
	}

	d.elements.Insert(element, position)
	return nil
}

func (d *document) GetElement(position int) (element, error) {
	if !d.positionIsValid(position) {
		return nil, errors.Errorf("Invalid element position `%d`", position)
	}

	return d.elements[position].(element), nil
}

func (d *document) getElements() vector.Vector {
	return d.elements
}

func (d *document) DeleteElement(position int) error {
	elem, err := d.GetElement(position)
	if err != nil {
		return errors.Wrap(err, "Failed to get element")
	}

	cmd := command.New(func() error {
		return d.DeleteElement(position)
	}, func() error {
		return d.insertElement(elem, position)
	})

	err = d.history.Record(cmd)
	return errors.Wrap(err, "Failed to record the command")
}

func (d *document) deleteElement(position int) error {
	if !d.positionIsValid(position) {
		return errors.Errorf("Invalid element position `%d`", position)
	}

	d.elements.Delete(position)
	return nil
}

func (d *document) positionIsValid(position int) bool {
	return 0 <= position && position < d.GetElementsCount()
}

func (d *document) String() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Title: %s", d.title))
	for index, elem := range d.elements {
		builder.WriteString(fmt.Sprintf("\n%d. %s", index, elem.(element).String()))
	}

	return builder.String()
}

func (d *document) acceptVisitor(v htmlFormatVisitor) string {
	return v.visitDocument(d)
}

func (d *document) Save(htmlPath string) error {
	dir := filepath.Dir(htmlPath)
	htmlStorage := storage.NewLocal(dir)

	err := storage.CopyAll(d.tempStorage, htmlStorage)
	if err != nil {
		return errors.Wrap(err, "Failed to copy files to HTML storage")
	}

	visitor := new(htmlFormatVisitorImpl)
	err = ioutil.WriteFile(htmlPath, []byte(d.acceptVisitor(visitor)), 0x666)
	return errors.Wrap(err, "Failed to write html file")
}

func (d *document) CanUndo() bool {
	return d.history.CanUndo()
}

func (d *document) CanRedo() bool {
	return d.history.CanRedo()
}

func (d *document) Undo() error {
	if !d.history.CanUndo() {
		return errors.New("Cannot undo")
	}

	err := d.history.Undo()
	return errors.Wrap(err, "Failed to undo")
}

func (d *document) Redo() error {
	if !d.history.CanRedo() {
		return errors.New("Cannot redo")
	}

	err := d.history.Redo()
	return errors.Wrap(err, "Failed to redo")
}

func (d *document) Release() error {
	err := d.tempStorage.Clear()
	if err != nil {
		return nil
	}

	return d.history.Release()
}
