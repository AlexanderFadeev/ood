package storage

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type local struct {
	root string
}

func NewLocal(rootDir string) Storage {
	return &local{
		root: rootDir,
	}
}

func (l *local) ListFiles() ([]string, error) {
	var files []string
	err := filepath.Walk(l.root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			path = strings.TrimPrefix(path, l.root)
			path = strings.TrimLeft(path, "/\\")
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to walk local storage root directory")
	}

	return files, nil
}

func (l *local) GetFile(key string) (io.ReadCloser, error) {
	if path.IsAbs(key) {
		return nil, errors.New("Key is not a relative path")
	}
	file, err := os.Open(l.getAbsolutePath(key))
	return file, errors.Wrap(err, "Failed to open file")
}

func (l *local) PutFile(key string, srcFile io.Reader) error {
	if path.IsAbs(key) {
		return errors.New("Key is not a relative path")
	}

	if l.fileExists(key) {
		return errors.New("File or directory with same name already exists")
	}

	dstFile, err := os.Create(l.getAbsolutePath(key))
	if err != nil {
		return errors.Wrap(err, "Failed to create file")
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return errors.Wrap(err, "Failed to copy file")
}

func (l *local) DeleteFile(key string) error {
	err := os.Remove(l.getAbsolutePath(key))
	return errors.Wrap(err, "Failed to remove file")
}

func (l *local) getAbsolutePath(key string) string {
	return path.Join(l.root, key)
}

func (l *local) fileExists(key string) bool {
	_, err := os.Stat(path.Join(l.root, key))
	return !os.IsNotExist(err)
}
