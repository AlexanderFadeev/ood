package document

import "fmt"

type Image interface {
	Element
	GetSize() (w, h int)
	SetSize(w, h int) error
}

type image struct {
	path   string
	width  int
	height int
}

func NewImage(path string, width, height int) Image {
	//TODO: add size check

	return &image{
		path:   path,
		width:  width,
		height: height,
	}
}

func (i *image) String() string {
	return fmt.Sprintf("%dx%d %s", i.width, i.height, i.path)
}

func (i *image) ToHTML() string {
	return fmt.Sprintf(`<img src="%s" width="%d" height="%d" />`, i.path, i.width, i.height)
}

func (i *image) GetSize() (w, h int) {
	return i.width, i.height
}

func (i *image) SetSize(w, h int) error {
	i.width, i.height = w, h
	return nil
}
