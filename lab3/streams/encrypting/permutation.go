package encrypting

import "math/rand"

type permutation []byte

type permutationMakerFunc func(seed int64, len uint) permutation

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
