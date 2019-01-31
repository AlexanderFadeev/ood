package presenter

import (
	"fmt"

	"ood/lab7/point"
	"ood/lab9/app/model"
	"ood/lab9/app/view"
)

type Presenter interface {
	Init()
}

type presenter struct {
	model       model.Model
	view        view.View
	lastShapeID int
}

func New(model model.Model, view view.View) Presenter {
	return &presenter{
		model: model,
		view:  view,
	}
}

func (p *presenter) Init() {
	p.view.DoOnShapeUpdate(func(id int, pos, size point.Vector) {
		fmt.Printf("Shape %d was moved/resized to position %v %v\n", id, pos, size)
	})

	p.view.DoOnButtonClick(p.view.AddRectangleButtonID(), func() {
		p.lastShapeID++
		p.view.AddRectangle(p.lastShapeID)
	})

	p.view.DoOnButtonClick(p.view.AddEllipseButtonID(), func() {
		p.lastShapeID++
		p.view.AddEllipse(p.lastShapeID)
	})

	p.view.RemoveLoader()
}
