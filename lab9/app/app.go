package app

import (
	"ood/lab9/app/model"
	"ood/lab9/app/presenter"
	"ood/lab9/app/view"
)

type App interface {
	Run()
}

type app struct {
	presenter presenter.Presenter
}

func New() App {
	m := model.New()
	v := view.New()
	p := presenter.New(m, v)
	return &app{
		presenter: p,
	}
}

func (a *app) Run() {
	a.presenter.Init()
	a.BlockGoroutine()
}

func (app) BlockGoroutine() {
	select {}
}
