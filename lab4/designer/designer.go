package designer

import (
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"ood/lab4/factory"
	"ood/lab4/picture_draft"
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
