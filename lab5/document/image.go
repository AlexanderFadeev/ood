package document

import (
	"fmt"
	"path"

	"github.com/AlexanderFadeev/ood/lab5/command"
	"github.com/AlexanderFadeev/ood/lab5/history"
	"github.com/AlexanderFadeev/ood/lab5/storage"

	"github.com/pkg/errors"
)

const maxImageSize = 5000

type image interface {
	element
	getPath() string
	getSize() (w, h int)
	setSize(w, h int) error
}

type imageImpl struct {
	path     string
	width    int
	height   int
	recorder history.Recorder
}

func (d *document) newImage(imgPath string, width, height int) (image, error) {
	if width <= 0 || maxImageSize < width {
		return nil, errors.New("Invalid image width")
	}
	if height <= 0 || maxImageSize < height {
		return nil, errors.New("Invalid image height")
	}

	ext := path.Ext(imgPath)
	pattern := fmt.Sprintf("images/%%d%s", ext)
	dstKey, err := storage.FindUnusedKey(d.tempStorage, pattern)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to find unused key")
	}

	file, err := d.localStorage.GetFile(imgPath)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to find file `%s`", imgPath)
	}
	defer file.Close()

	err = d.tempStorage.PutFile(*dstKey, file)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to copy file to temp storage")
	}

	return &imageImpl{
		path:     *dstKey,
		width:    width,
		height:   height,
		recorder: d.history,
	}, nil
}

func (i *imageImpl) String() string {
	return fmt.Sprintf("%dx%d %s", i.width, i.height, i.path)
}

func (i *imageImpl) acceptVisitor(visitor htmlFormatVisitor) string {
	return visitor.visitImage(i)
}

func (i *imageImpl) getPath() string {
	return i.path
}

func (i *imageImpl) getSize() (w, h int) {
	return i.width, i.height
}

func (i *imageImpl) setSize(width, height int) error {
	oldWidth, oldHeight := i.getSize()

	cmd := command.New(func() error {
		i.width, i.height = width, height
		return nil
	}, func() error {
		i.width, i.height = oldWidth, oldHeight
		return nil
	})

	err := i.recorder.Record(cmd)
	return errors.Wrap(err, "Failed to record the command")
}
