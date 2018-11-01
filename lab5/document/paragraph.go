package document

import (
	"fmt"

	"ood/lab5/command"
	"ood/lab5/history"

	"github.com/pkg/errors"
)

type paragraph interface {
	element
	getText() string
	setText(string) error
}

type paragraphImpl struct {
	text     string
	recorder history.Recorder
}

func newParagraph(text string, recorder history.Recorder) paragraph {
	return &paragraphImpl{
		text:     text,
		recorder: recorder,
	}
}

func (p *paragraphImpl) String() string {
	return fmt.Sprintf("Paragraph: %s", p.text)
}

func (p *paragraphImpl) acceptVisitor(visitor htmlFormatVisitor) string {
	return visitor.visitParagraph(p)
}

func (p *paragraphImpl) getText() string {
	return p.text
}

func (p *paragraphImpl) setText(text string) error {
	oldText := p.text

	cmd := command.New(func() error {
		p.text = text
		return nil
	}, func() error {
		p.text = oldText
		return nil
	})

	err := p.recorder.Record(cmd)
	return errors.Wrap(err, "Failed to record the command")
}
