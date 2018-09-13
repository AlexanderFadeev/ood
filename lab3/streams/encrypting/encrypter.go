package encrypting

type Encrypter interface {
	Encrypt([]byte) []byte
}

type encrypter struct {
	perm permutation
}

func newEncrypter(seed int64, makerFunc permutationMakerFunc) Encrypter {
	return &encrypter{
		perm: makerFunc(seed, 1<<8),
	}
}

func NewEncrypter(seed int64) Encrypter {
	return newEncrypter(seed, newPermutation)
}

func (e *encrypter) Encrypt(data []byte) []byte {
	encryptedData := make([]byte, len(data))

	for index, value := range data {
		encryptedData[index] = e.perm[value]
	}

	return encryptedData
}
