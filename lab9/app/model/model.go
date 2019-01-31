package model

type Model interface {
}

type model struct {
}

func New() Model {
	return &model{}
}
