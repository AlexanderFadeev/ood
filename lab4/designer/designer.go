package designer

import (
	"bufio"
	"github.com/AlexanderFadeev/ood/lab4/factory"
	"github.com/AlexanderFadeev/ood/lab4/picture_draft"
	"github.com/pkg/errors"
	"io"
)

type Designer interface {
	CreateDraft(io.Reader) (picture_draft.PictureDraft, error)
}

type designer struct{}

func New() Designer {
	return new(designer)
}

func (d *designer) CreateDraft(reader io.Reader) (picture_draft.PictureDraft, error) {
	f := factory.New()
	var draft picture_draft.PictureDraft

	s := bufio.NewScanner(reader)
	for s.Scan() {
		shape, err := f.CreateShape(s.Text())
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to create shape from desc `%s`", s.Text())
		}

		draft = append(draft, shape)
	}
	if err := s.Err(); err != nil {
		return nil, errors.Wrap(err, "Failed to scan data")
	}

	return draft, nil
}
