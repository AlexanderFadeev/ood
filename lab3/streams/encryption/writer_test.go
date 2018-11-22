package encryption

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentityWriter(t *testing.T) {
	var buf bytes.Buffer
	encrypter := newCrypter(newIdentityPermutation(byteValuesCount))
	writer := newEncodingWriter(&buf, encrypter)

	writer.Write(getSampleData())
	result := buf.Bytes()

	assert.Equal(t, getSampleData(), result)
}

func TestCaesarWriter(t *testing.T) {
	var buf bytes.Buffer
	encrypter := newCrypter(newCaesarPermutation(byteValuesCount, sampleShift))
	writer := newEncodingWriter(&buf, encrypter)

	writer.Write(getSampleData())
	result := buf.Bytes()

	expected := caesarEncode(getSampleData(), sampleShift)
	assert.Equal(t, expected, result)
}
