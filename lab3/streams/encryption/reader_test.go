package encryption

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentityReader(t *testing.T) {
	var buf bytes.Buffer
	buf.Write(getSampleData())

	decrypter := newCrypter(newIdentityPermutation(byteValuesCount))
	reader := newEncodingReader(&buf, decrypter)

	data, err := ioutil.ReadAll(reader)
	assert.Nil(t, err)

	assert.Equal(t, getSampleData(), data)
}

func TestCaesarReader(t *testing.T) {
	var buf bytes.Buffer
	buf.Write(caesarEncode(getSampleData(), sampleShift))

	decrypter := newCrypter(newCaesarPermutation(byteValuesCount, -sampleShift))
	reader := newEncodingReader(&buf, decrypter)

	data, err := ioutil.ReadAll(reader)
	assert.Nil(t, err)

	assert.Equal(t, getSampleData(), data)
}
