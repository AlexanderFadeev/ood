package decorator

import (
	"io"
	"ood/lab3/streams/compression"
	"ood/lab3/streams/encryption"
)

type WriterDecorator interface {
	DecorateWriter(io.WriteCloser) io.WriteCloser
}

type CompositeWriterDecorator []WriterDecorator

func (c CompositeWriterDecorator) DecorateWriter(w io.WriteCloser) io.WriteCloser {
	result := w
	for _, decorator := range c {
		result = decorator.DecorateWriter(result)
	}
	return result
}

type CompressionWriterDecorator struct{}

func (CompressionWriterDecorator) DecorateWriter(w io.WriteCloser) io.WriteCloser {
	return compression.NewCompressor(w)
}

type encryptionWriterDecorator struct {
	key int64
}

func NewEncryptionWriterDecorator(key int64) WriterDecorator {
	return &encryptionWriterDecorator{
		key: key,
	}
}

func (e *encryptionWriterDecorator) DecorateWriter(w io.WriteCloser) io.WriteCloser {
	return encryption.NewEncryptingWriter(w, e.key)
}
