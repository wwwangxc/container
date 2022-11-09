package hashset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_Add(t *testing.T) {
	set := New[int]()
	set.Add(1, 2, 3)
	assert.Equal(t, 3, set.Size())
	assert.True(t, set.Contains(1))
	assert.True(t, set.Contains(2))
	assert.True(t, set.Contains(3))
	assert.False(t, set.Contains(5))
}

func TestSet_Remove(t *testing.T) {
	set := New(1, 2, 3)
	assert.Equal(t, 3, set.Size())
	assert.True(t, set.Contains(1, 2, 3))
	set.Remove(1, 2, 3)
	assert.Equal(t, 0, set.Size())
	assert.False(t, set.Contains(1, 2, 3))
}

func TestSet_Contains(t *testing.T) {
	set := New(1, 2, 3)
	assert.True(t, set.Contains(1))
	assert.True(t, set.Contains(2))
	assert.True(t, set.Contains(3))
	assert.True(t, set.Contains(1, 2))
	assert.True(t, set.Contains(1, 3))
	assert.True(t, set.Contains(2, 3))
	assert.True(t, set.Contains(1, 2, 3))
	assert.False(t, set.Contains(4))
	assert.False(t, set.Contains(1, 2, 3, 4))
}

func TestSet_Union(t *testing.T) {
	set := New(1, 2, 3)
	target := New(4, 5, 6)

	newSet := set.Union(target)
	assert.Equal(t, set.Size()+target.Size(), newSet.Size())

	for _, v := range set.Values() {
		assert.True(t, newSet.Contains(v))
	}

	for _, v := range target.Values() {
		assert.True(t, newSet.Contains(v))
	}
}

func TestSet_Intersection(t *testing.T) {
	set := New(1, 2, 3)
	target := New(1, 2, 4, 5, 6)

	newSet := set.Intersection(target)
	for _, v := range newSet.Values() {
		assert.True(t, set.Contains(v) && target.Contains(v))
	}
}

func TestSet_Difference(t *testing.T) {
	set := New(1, 2, 3)
	target := New(1, 2, 4, 5, 6)

	newSet := set.Difference(target)
	for _, v := range newSet.Values() {
		assert.True(t, set.Contains(v) && !target.Contains(v))
	}
}
