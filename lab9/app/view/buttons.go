package view

import "syscall/js"

type ButtonsIDProvider interface {
	AddRectangleButtonID() int
	AddEllipseButtonID() int
}

type buttonsIDProvider struct{}

func (buttonsIDProvider) AddRectangleButtonID() int {
	return js.Global().Get("buttonAddRectangle").Int()
}

func (buttonsIDProvider) AddEllipseButtonID() int {
	return js.Global().Get("buttonAddEllipse").Int()
}
