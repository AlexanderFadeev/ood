package encryption

func newIdentityPermutation(len uint) permutation {
	return newCaesarPermutation(len, 0)
}

func newCaesarPermutation(len, shift uint) permutation {
	perm := make([]byte, len)

	for index := uint(0); index < len; index++ {
		perm[index] = byte((index + shift) % (1 << 8))
	}

	return perm
}
