package decorator

import (
	"io"

	"github.com/AlexanderFadeev/ood/lab3/streams/compression"
	"github.com/AlexanderFadeev/ood/lab3/streams/encryption"
)

type ReaderDecorator interface {
	DecorateReader(io.ReadCloser) io.ReadCloser
}

type CompositeReaderDecorator []ReaderDecorator

func (c CompositeReaderDecorator) DecorateReader(r io.ReadCloser) io.ReadCloser {
	result := r
	for _, decorator := range c {
		result = decorator.DecorateReader(result)
	}
	return result
}

type DecompressionReaderDecorator struct{}

func (DecompressionReaderDecorator) DecorateReader(r io.ReadCloser) io.ReadCloser {
	return compression.NewDecompressor(r)
}

type decryptionReaderDecorator struct {
	key int64
}

func NewDecryptionReaderDecorator(key int64) ReaderDecorator {
	return &decryptionReaderDecorator{
		key: key,
	}
}

func (d *decryptionReaderDecorator) DecorateReader(r io.ReadCloser) io.ReadCloser {
	return encryption.NewDecryptingReader(r, d.key)
}
