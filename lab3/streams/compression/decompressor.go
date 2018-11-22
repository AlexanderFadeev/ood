package compression

import (
	"compress/gzip"
	"io"
)

type decompressor struct {
	io.ReadCloser
	origReader io.ReadCloser
}

func NewDecompressor(r io.ReadCloser) io.ReadCloser {
	gzipReader, err := gzip.NewReader(r)
	if err != nil {
		panic(err)
	}

	return &decompressor{
		ReadCloser: gzipReader,
		origReader: r,
	}
}

func (c *decompressor) Close() error {
	err1 := c.ReadCloser.Close()
	err2 := c.origReader.Close()
	if err1 != nil {
		return err1
	}
	return err2
}
