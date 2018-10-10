package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorPush(t *testing.T) {
	vec := make(Vector, 0)

	count := 42
	for i := 0; i < 42; i++ {
		vec.Push(i)
	}

	assert.Equal(t, count, len(vec))

	for i := 0; i < 42; i++ {
		assert.Equal(t, i, vec[i])
	}
}

func TestVectorInsert(t *testing.T) {
	vec := Vector{1, 2, 4, 5}

	vec.Insert(3, 2)

	assert.Equal(t, Vector{1, 2, 3, 4, 5}, vec)
}

func TestVectorDelete(t *testing.T) {
	vec := Vector{1, 2, 3, 4, 5}

	vec.Delete(2)

	assert.Equal(t, Vector{1, 2, 4, 5}, vec)
}
