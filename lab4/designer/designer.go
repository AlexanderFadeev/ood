package designer

import (
	"github.com/AlexanderFadeev/ood/lab4/factory"
	"github.com/AlexanderFadeev/ood/lab4/picture_draft"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"strings"
)

type Designer interface {
	CreateDraft(io.Reader) (picture_draft.PictureDraft, error)
}

type designer struct{}

func New() Designer {
	return new(designer)
}

func (d *designer) CreateDraft(reader io.Reader) (picture_draft.PictureDraft, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read from reader")
	}

	lines := strings.Split(string(data), "\n")
	draft := make(picture_draft.PictureDraft, len(lines))
	f := factory.New()

	for index, line := range lines {
		draft[index], err = f.CreateShape(line)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to create shape from desc `%s`", line)
		}
	}

	return draft, nil
}
