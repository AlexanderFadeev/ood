package editor

import (
	"fmt"
	"path"

	"ood/lab5/command"
	"ood/lab5/document"
	"ood/lab5/storage"

	"github.com/pkg/errors"
)

func (e *editor) newInsertImageCommand(pos int, width, height int, imgPath string) (command.Command, error) {
	ext := path.Ext(imgPath)
	pattern := fmt.Sprintf("images/%%d%s", ext)
	dstKey, err := storage.FindUnusedKey(e.tempStorage, pattern)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to find unused key")
	}

	file, err := storage.FindInAny(imgPath, e.workDirStorage, e.globalStorage)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to find file `%s`", imgPath)
	}
	defer file.Close()

	err = e.tempStorage.PutFile(*dstKey, file)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to copy file to temp storage")
	}

	img := document.NewImage(*dstKey, width, height)

	cmd := command.NewWithRelease(func() error {
		return e.doc.InsertElement(img, pos)
	}, func() error {
		return e.doc.DeleteElement(pos)
	}, func() error {
		return e.tempStorage.DeleteFile(*dstKey)
	})

	return cmd, nil
}
