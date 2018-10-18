package file_system

import (
	"io"
	"os"
	"path/filepath"
)

type FileSystem interface {
	Create(path string) (io.WriteCloser, error)
	Open(path string) (io.ReadCloser, error)
	Remove(path string) error
	Exists(path string) (bool, error)
	MkdirAll(path string) error
	GetWorkDir() (string, error)
	Walk(root string, walkFn filepath.WalkFunc) error
}

type fileSystem struct{}

func New() FileSystem {
	return new(fileSystem)
}

func (fs *fileSystem) Open(path string) (io.ReadCloser, error) {
	return os.Open(path)
}

func (fs *fileSystem) Create(path string) (io.WriteCloser, error) {
	return os.Create(path)
}

func (fs *fileSystem) Remove(path string) error {
	return os.Remove(path)
}

func (fs *fileSystem) Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (fs *fileSystem) GetWorkDir() (string, error) {
	return os.Getwd()
}

func (fs *fileSystem) MkdirAll(path string) error {
	return os.MkdirAll(path, 0666)
}

func (fs *fileSystem) Walk(root string, walkFn filepath.WalkFunc) error {
	return filepath.Walk(root, walkFn)
}
