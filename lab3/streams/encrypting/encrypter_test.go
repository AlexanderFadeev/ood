package encrypting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentityEncrypter(t *testing.T) {
	encrypter := newEncrypter(sampleSeed, newIdentityPermutation)
	result := encrypter.Encrypt(sampleData)
	assert.Equal(t, sampleData, result)
}

func TestEncryptedDataLength(t *testing.T) {
	encrypter := NewEncrypter(sampleSeed)
	result := encrypter.Encrypt(sampleData)
	assert.Equal(t, len(sampleData), len(result))
}

func TestEncryption(t *testing.T) {
	expected := caesarEncode(sampleData, sampleShift)
	encrypter := newEncrypter(sampleSeed, newCaesarPermutationMakerFunc(sampleShift))
	result := encrypter.Encrypt(sampleData)
	assert.Equal(t, expected, result)
}
