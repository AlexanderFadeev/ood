package factory

import (
	"github.com/AlexanderFadeev/ood/lab4/shape"

	"github.com/pkg/errors"
)

type Factory interface {
	CreateShape(description string) (shape.Shape, error)
}

type factory struct{}

func New() Factory {
	return new(factory)
}

func (f *factory) CreateShape(descStr string) (shape.Shape, error) {
	desc, err := newDescription(descStr)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create new shape description")
	}

	return desc.toShape()
}
