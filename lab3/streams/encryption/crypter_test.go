package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentityCrypter(t *testing.T) {
	data := getSampleData()
	expected := getSampleData()

	crypter := newCrypter(newIdentityPermutation(byteValuesCount))

	crypter.crypt(data)
	assert.Equal(t, expected, data)
}

func TestCaesarCrypter(t *testing.T) {
	data := getSampleData()
	expected := caesarEncode(data, sampleShift)

	crypter := newCrypter(newCaesarPermutation(byteValuesCount, sampleShift))

	crypter.crypt(data)
	assert.Equal(t, expected, data)
}

func TestCryptedDataLength(t *testing.T) {
	data := getSampleData()
	expectedLen := len(data)

	encrypter := newEncrypter(sampleSeed)

	encrypter.crypt(data)
	assert.Equal(t, expectedLen, len(data))
}
