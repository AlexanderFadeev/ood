package storage

import (
	"io"
	"os"
	"path"
	"strings"

	"ood/lab5/file_system"

	"github.com/pkg/errors"
)

type local struct {
	root string
	fs   file_system.FileSystem
}

func newLocal(rootDir string, fs file_system.FileSystem) (Storage, error) {
	if !path.IsAbs(rootDir) {
		wd, err := fs.GetWorkDir()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to get workdir")
		}

		rootDir = path.Join(wd, rootDir)
	}

	return &local{
		root: rootDir,
		fs:   fs,
	}, nil
}

func NewLocal(rootDir string) (Storage, error) {
	return newLocal(rootDir, file_system.New())
}

func (l *local) ListFiles() ([]string, error) {
	var files []string
	err := l.fs.Walk(l.root, func(path string, info os.FileInfo, err error) error {
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

func (l *local) HasFile(key string) (bool, error) {
	exists, err := l.fs.Exists(path.Join(l.root, key))
	return exists, errors.Wrap(err, "Failed to check if file exists if file system")
}

func (l *local) GetFile(key string) (io.ReadCloser, error) {
	file, err := l.fs.Open(l.getAbsolutePath(key))
	return file, errors.Wrap(err, "Failed to open file")
}

func (l *local) PutFile(key string, srcFile io.Reader) error {
	exists, err := l.fs.Exists(l.getAbsolutePath(key))
	if err != nil {
		return errors.Wrap(err, "Failed to check if file exists in file system")
	}
	if exists {
		return errors.Wrap(err, "File or directory with same name already exists")
	}

	err = l.fs.MkdirAll(path.Dir(l.getAbsolutePath(key)))
	if err != nil {
		return errors.Wrap(err, "Failed to make dirs for file")
	}

	dstFile, err := l.fs.Create(l.getAbsolutePath(key))
	if err != nil {
		return errors.Wrap(err, "Failed to create file")
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return errors.Wrap(err, "Failed to copy file")
}

func (l *local) DeleteFile(key string) error {
	err := l.fs.Remove(l.getAbsolutePath(key))
	return errors.Wrap(err, "Failed to remove file")
}

func (l *local) getAbsolutePath(key string) string {
	return path.Join(l.root, key)
}
