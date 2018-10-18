package document

import "fmt"

type Paragraph interface {
	Element
	GetText() string
	SetText(string)
}

type paragraph struct {
	text string
}

func NewParagraph(text string) Paragraph {
	return &paragraph{
		text: text,
	}
}

func (p *paragraph) String() string {
	return fmt.Sprintf("Paragraph: %s", p.text)
}

func (p *paragraph) ToHTML() string {
	return fmt.Sprintf("<p>%s</p>", p.text)
}

func (p *paragraph) GetText() string {
	return p.text
}

func (p *paragraph) SetText(text string) {
	p.text = text
}
