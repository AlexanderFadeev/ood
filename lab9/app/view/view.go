package view

import (
	"ood/lab2/signal"
	"ood/lab7/point"
	"syscall/js"
)

type View interface {
	ButtonsIDProvider

	AddRectangle(id int)
	AddEllipse(id int)

	DoOnShapeUpdate(ShapeUpdateHandler)
	DoOnButtonClick(int, VoidHandler)

	RemoveLoader()
}

type view struct {
	buttonsIDProvider
	signal signal.Signal
	jsImpl js.Value
}

func New() View {
	return &view{
		jsImpl: js.Global().Get("View").New(),
	}
}

func (v view) RemoveLoader() {
	v.jsImpl.Call("removeLoader")
}

func (v *view) AddRectangle(id int) {
	v.jsImpl.Call("addRectangle", id)
}

func (v *view) AddEllipse(id int) {
	v.jsImpl.Call("addEllipse", id)
}

type ShapeUpdateHandler func(id int, pos, size point.Vector)

func (v *view) DoOnShapeUpdate(handler ShapeUpdateHandler) {
	cb := js.NewCallback(func(args []js.Value) {
		handler(
			args[0].Int(),
			point.Vector{
				args[1].Float(),
				args[2].Float(),
			},
			point.Vector{
				args[3].Float(),
				args[4].Float(),
			},
		)
	})

	v.jsImpl.Call("doOnShapeUpdate", cb)
}

type VoidHandler func()

func (v *view) DoOnButtonClick(id int, handler VoidHandler) {
	cb := js.NewCallback(func(args []js.Value) {
		handler()
	})

	v.jsImpl.Call("doOnButtonClick", id, cb)
}
