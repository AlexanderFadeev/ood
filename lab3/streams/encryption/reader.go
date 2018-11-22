package encryption

import (
	"io"
)

type decryptingReader struct {
	reader  io.ReadCloser
	crypter crypter
}

func newEncodingReader(reader io.ReadCloser, crypter crypter) io.ReadCloser {
	return &decryptingReader{
		reader:  reader,
		crypter: crypter,
	}
}

func NewDecryptingReader(reader io.ReadCloser, key int64) io.ReadCloser {
	return newEncodingReader(reader, newDecrypter(key))
}

func (r *decryptingReader) Read(data []byte) (int, error) {
	n, err := r.reader.Read(data)
	if err != nil {
		return n, err
	}

	r.crypter.crypt(data[:n])
	return n, nil
}

func (r *decryptingReader) Close() error {
	return r.reader.Close()
}
