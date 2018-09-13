package encrypting

func newIdentityPermutation(_ int64, len uint) permutation {
	return newCaesarPermutationMakerFunc(0)(0, len)
}

func newCaesarPermutationMakerFunc(shift uint) permutationMakerFunc {
	return func(_ int64, len uint) permutation {
		perm := make([]byte, len)

		for index := uint(0); index < len; index++ {
			perm[index] = byte((index + shift) % (1 << 8))
		}

		return perm
	}
}
