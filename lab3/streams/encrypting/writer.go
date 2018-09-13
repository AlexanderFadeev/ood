package encrypting

import "io"

type encryptingWriter struct {
	writer    io.Writer
	encrypter Encrypter
}

func NewEncodingWriter(writer io.Writer, encrypter Encrypter) io.Writer {
	return &encryptingWriter{
		writer:    writer,
		encrypter: encrypter,
	}
}

func (e *encryptingWriter) Write(data []byte) (int, error) {
	return e.writer.Write(e.encrypter.Encrypt(data))
}
