package doublylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Any(t *testing.T) {
	list := New(1, 2, 3)

	got := list.Any(func(_, value int) bool {
		return value == 1
	})
	assert.True(t, got)

	got = list.Any(func(_, value int) bool {
		return value == 0
	})
	assert.False(t, got)
}

func BenchmarkList_Any(b *testing.B) {
	values := make([]int, 0, b.N)
	for i := 0; i < 1000; i++ {
		values = append(values, i)
	}

	list := New(values...)
	f := func(index, value int) bool {
		return false
	}

	for i := 0; i < b.N; i++ {
		_ = list.Any(f)
	}
}

func TestList_All(t *testing.T) {
	list := New(1, 2, 3)

	got := list.All(func(_, value int) bool {
		return value == 1
	})
	assert.False(t, got)

	got = list.All(func(_, value int) bool {
		return false
	})
	assert.False(t, got)

	got = list.All(func(_, value int) bool {
		return true
	})
	assert.True(t, got)
}

func BenchmarkList_All(b *testing.B) {
	values := make([]int, 0, b.N)
	for i := 0; i < 1000; i++ {
		values = append(values, i)
	}

	list := New(values...)
	f := func(index, value int) bool {
		return true
	}

	for i := 0; i < b.N; i++ {
		_ = list.All(f)
	}
}

func TestList_Find(t *testing.T) {
	list := New(1, 2, 3, 1)

	index, value, exist := list.Find(func(index, value int) bool {
		return value == 1
	})
	assert.True(t, exist)
	assert.Equal(t, 0, index)
	assert.Equal(t, 1, value)

	index, value, exist = list.Find(func(index, value int) bool {
		return value == 0
	})
	assert.False(t, exist)
	assert.Equal(t, -1, index)
	assert.Equal(t, 0, value)
}

func BenchmarkList_Find(b *testing.B) {
	values := make([]int, 0, b.N)
	for i := 0; i < 1000; i++ {
		values = append(values, i)
	}

	list := New(values...)
	f := func(index, value int) bool {
		return index%10 == 0
	}

	for i := 0; i < b.N; i++ {
		_, _, _ = list.Find(f)
	}
}

func TestList_Select(t *testing.T) {
	list := New(1, 2, 3)

	values := list.Select(func(index, value int) bool {
		return value > 1
	})
	assert.Equal(t, []int{2, 3}, values)

	values = list.Select(func(index, value int) bool {
		return value == 0
	})
	assert.Equal(t, []int(nil), values)
}

func BenchmarkList_Select(b *testing.B) {
	values := make([]int, 0, b.N)
	for i := 0; i < 1000; i++ {
		values = append(values, i)
	}

	list := New(values...)
	f := func(index, value int) bool {
		return index%10 == 0
	}

	for i := 0; i < b.N; i++ {
		_ = list.Select(f)
	}
}
