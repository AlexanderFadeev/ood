package encrypting

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentityWriter(t *testing.T) {
	var buf bytes.Buffer
	encrypter := newEncrypter(sampleSeed, newIdentityPermutation)
	writer := NewEncodingWriter(&buf, encrypter)

	writer.Write(sampleData)
	result := buf.Bytes()

	assert.Equal(t, sampleData, result)
}

func TestWriter(t *testing.T) {
	var buf bytes.Buffer
	encrypter := newEncrypter(sampleSeed, newCaesarPermutationMakerFunc(sampleShift))
	writer := NewEncodingWriter(&buf, encrypter)

	writer.Write(sampleData)
	result := buf.Bytes()

	expected := caesarEncode(sampleData, sampleShift)
	assert.Equal(t, expected, result)
}
