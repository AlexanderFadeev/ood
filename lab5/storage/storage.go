package storage

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

type Storage interface {
	ListFiles() ([]string, error)
	HasFile(key string) (bool, error)
	GetFile(key string) (io.ReadCloser, error)
	PutFile(key string, file io.Reader) error
	DeleteFile(key string) error
}

func Copy(srcStorage, dstStorage Storage, srcKey, dstKey string) error {
	file, err := srcStorage.GetFile(srcKey)
	if err != nil {
		return errors.Wrap(err, "Failed to get file from source storage")
	}

	err = dstStorage.PutFile(dstKey, file)
	if err != nil {
		return errors.Wrap(err, "Failed to copy file")
	}

	err = file.Close()
	return errors.Wrap(err, "Failed to close src file")
}

func CopyAll(src, dst Storage) error {
	keys, err := src.ListFiles()
	if err != nil {
		return errors.Wrap(err, "Failed to get list of files from src storage")
	}

	for _, key := range keys {
		file, err := src.GetFile(key)
		if err != nil {
			return errors.Wrap(err, "Failed to get file from src storage")
		}

		err = dst.PutFile(key, file)
		if err != nil {
			return errors.Wrap(err, "Failed to put file in dst storage")
		}
	}

	return nil
}

func FindUnusedKey(storage Storage, pattern string) (*string, error) {
	list, err := storage.ListFiles()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to list files")
	}

	id := 0
	key := fmt.Sprintf(pattern, id)
	for keyIsInList(key, list) {
		id++
		key = fmt.Sprintf(pattern, id)
	}

	return &key, nil
}

func FindInAny(key string, storages ...Storage) (io.ReadCloser, error) {
	for _, storage := range storages {
		exists, err := storage.HasFile(key)
		if !exists || err != nil {
			continue
		}

		file, err := storage.GetFile(key)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to get file `%s` from storage", key)
		}

		return file, nil
	}

	return nil, errors.Errorf("File `%s` not found in any of storages", key)
}

func keyIsInList(key string, list []string) bool {
	for _, keyFromList := range list {
		if key == keyFromList {
			return true
		}
	}
	return false
}
