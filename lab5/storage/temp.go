package storage

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type TempStorage interface {
	Storage

	Clear() error
}

type temp struct {
	local
}

func NewTemp() (TempStorage, error) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create temp dir")
	}

	return &temp{
		local: local{
			root: tempDir,
		},
	}, nil
}

func (t *temp) Clear() error {
	err := os.RemoveAll(t.local.root)
	return errors.Wrap(err, "Failed to remove temp dir")
}
