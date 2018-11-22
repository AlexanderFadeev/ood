package compression

import (
	"compress/gzip"
	"io"
)

type compressor struct {
	io.WriteCloser
	origWriter io.WriteCloser
}

func NewCompressor(w io.WriteCloser) io.WriteCloser {
	return &compressor{
		WriteCloser: gzip.NewWriter(w),
		origWriter:  w,
	}
}

func (c *compressor) Close() error {
	err1 := c.WriteCloser.Close()
	err2 := c.origWriter.Close()
	if err1 != nil {
		return err1
	}
	return err2
}
