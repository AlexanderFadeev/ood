package document

import (
	"fmt"
	"github.com/pkg/errors"
	"ood/lab5/vector"
	"strings"
)

type Document interface {
	fmt.Stringer

	GetElementsCount() int
	GetTitle() string
	SetTitle(string)
	InsertElement(element Element, position int) error
	GetElement(position int) (Element, error)
	DeleteElement(position int) error
	ToHTML() string
}

type document struct {
	title    string
	elements vector.Vector
}

func New() Document {
	return new(document)
}

func (d *document) GetElementsCount() int {
	return len(d.elements)
}

func (d *document) GetTitle() string {
	return d.title
}

func (d *document) SetTitle(title string) {
	d.title = title
}

func (d *document) InsertElement(element Element, position int) error {
	if !d.positionIsValid(position) && position != d.GetElementsCount() {
		return errors.Errorf("Invalid element position `%d`", position)
	}

	d.elements.Insert(element, position)
	return nil
}

func (d *document) GetElement(position int) (Element, error) {
	if !d.positionIsValid(position) {
		return nil, errors.Errorf("Invalid element position `%d`", position)
	}

	return d.elements[position].(Element), nil
}

func (d *document) DeleteElement(position int) error {
	if !d.positionIsValid(position) {
		return errors.Errorf("Invalid element position `%d`", position)
	}

	d.elements.Delete(position)
	return nil
}

func (d *document) String() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Title: %s", d.title))
	for index, element := range d.elements {
		builder.WriteString(fmt.Sprintf("\n%d. %s", index, element.(Element).String()))
	}

	return builder.String()
}

func (d *document) ToHTML() string {
	const htmlTemplate = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>%s</title>
</head>
<body>%s
</body>
</html>
`

	var builder strings.Builder
	for _, element := range d.elements {
		builder.WriteString(fmt.Sprintf("\n%s", element.(Element).ToHTML()))
	}

	return fmt.Sprintf(htmlTemplate, d.GetTitle(), builder.String())
}

func (d *document) positionIsValid(position int) bool {
	return 0 <= position && position < d.GetElementsCount()
}
