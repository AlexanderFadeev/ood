package encryption

import "math/rand"

type permutation []byte

func newPermutation(seed int64, len uint) permutation {
	src := rand.NewSource(seed)
	rnd := rand.New(src)

	intPerm := rnd.Perm(int(len))
	perm := make([]byte, len)

	for index, value := range intPerm {
		perm[index] = byte(value)
	}

	return perm
}

func (p permutation) inverted() permutation {
	result := make(permutation, len(p))
	for index, value := range p {
		result[value] = byte(index)
	}
	return result
}
