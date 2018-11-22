package encryption

import "io"

type encryptingWriter struct {
	writer  io.WriteCloser
	crypter crypter
}

func newEncodingWriter(writer io.WriteCloser, encrypter crypter) io.WriteCloser {
	return &encryptingWriter{
		writer:  writer,
		crypter: encrypter,
	}
}

func NewEncryptingWriter(writer io.WriteCloser, key int64) io.WriteCloser {
	return newEncodingWriter(writer, newEncrypter(key))
}

func (w *encryptingWriter) Write(data []byte) (int, error) {
	dataCopy := make([]byte, len(data))
	copy(dataCopy, data)
	w.crypter.crypt(dataCopy)
	return w.writer.Write(dataCopy)
}

func (w *encryptingWriter) Close() error {
	return w.writer.Close()
}
