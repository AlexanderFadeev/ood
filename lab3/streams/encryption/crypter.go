package encryption

const (
	byteSize        = 8
	byteValuesCount = 1 << byteSize
)

type crypter interface {
	crypt(data []byte)
}

type crypterImpl struct {
	perm permutation
}

func newCrypter(perm permutation) crypter {
	return &crypterImpl{
		perm: perm,
	}
}

func (e *crypterImpl) crypt(data []byte) {
	for index, value := range data {
		data[index] = e.perm[value]
	}
}

func newEncrypter(seed int64) crypter {
	return newCrypter(newPermutation(seed, byteValuesCount))
}

func newDecrypter(seed int64) crypter {
	return newCrypter(newPermutation(seed, byteValuesCount).inverted())
}
