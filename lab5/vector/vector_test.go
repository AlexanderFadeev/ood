package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorPush(t *testing.T) {
	vec := make(Vector, 0)

	const count = 42
	for i := 0; i < count; i++ {
		vec.Push(i)
	}

	assert.Equal(t, count, len(vec))

	for i := 0; i < count; i++ {
		assert.Equal(t, i, vec[i])
	}
}

func TestVectorInsert(t *testing.T) {
	vec := Vector{1, 2, 4, 5}

	vec.Insert(3, 2)

	assert.Equal(t, Vector{1, 2, 3, 4, 5}, vec)
}

func TestVectorInsertFront(t *testing.T) {
	vec := Vector{2, 3, 4, 5}

	vec.Insert(1, 0)

	assert.Equal(t, Vector{1, 2, 3, 4, 5}, vec)
}

func TestVectorInsertBack(t *testing.T) {
	vec := Vector{1, 2, 3, 4}

	vec.Insert(5, 4)

	assert.Equal(t, Vector{1, 2, 3, 4, 5}, vec)
}

func TestVectorDelete(t *testing.T) {
	vec := Vector{1, 2, 3, 4, 5}

	vec.Delete(2)

	assert.Equal(t, Vector{1, 2, 4, 5}, vec)
}

func TestVectorDeleteFirst(t *testing.T) {
	vec := Vector{1, 2, 3, 4, 5}

	vec.Delete(0)

	assert.Equal(t, Vector{2, 3, 4, 5}, vec)
}

func TestVectorDeleteTheOnlyElement(t *testing.T) {
	vec := Vector{1}

	vec.Delete(0)

	assert.Equal(t, Vector{}, vec)
}

func TestVectorDeleteBack(t *testing.T) {
	vec := Vector{1, 2, 3, 4, 5}

	vec.Delete(4)

	assert.Equal(t, Vector{1, 2, 3, 4}, vec)
}

func TestVectorPop(t *testing.T) {
	vec := Vector{1, 2, 3, 4, 5}

	val := vec.Pop()
	assert.Equal(t, 5, val)

	assert.Equal(t, Vector{1, 2, 3, 4}, vec)
}

func TestVectorPopTheOnlyElement(t *testing.T) {
	vec := Vector{1}

	val := vec.Pop()
	assert.Equal(t, 1, val)

	assert.Equal(t, Vector{}, vec)
}
