package storage

import "io"

type Storage interface {
	ListFiles() ([]string, error)
	GetFile(key string) (io.ReadCloser, error)
	PutFile(key string, file io.Reader) error
	DeleteFile(key string) error
}
