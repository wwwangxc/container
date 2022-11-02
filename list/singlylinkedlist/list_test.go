package singlylinkedlist

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestList_Get(t *testing.T) {
	list := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	got, bool := list.Get(9)
	assert.True(t, bool)
	assert.Equal(t, 9, got)

	got, bool = list.Get(11)
	assert.False(t, bool)
	assert.Equal(t, 0, got)
}

func BenchmarkList_Get(b *testing.B) {
	list := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < b.N; i++ {
		_, _ = list.Get(i)
	}
}

func TestList_Contains(t *testing.T) {
	intValues := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	list := New(intValues...)
	assert.False(t, list.Contains(111, 222, 333))
	assert.True(t, list.Contains(intValues...))
	for _, v := range intValues {
		assert.True(t, list.Contains(v))
	}
}

func BenchmarkList_Contains(b *testing.B) {
	list := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < b.N; i++ {
		_ = list.Contains(i)
	}
}

func TestList_IndexOf(t *testing.T) {
	intValues := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	list := New(intValues...)
	assert.Equal(t, -1, list.IndexOf(111))
	for i, v := range intValues {
		assert.Equal(t, i, list.IndexOf(v))
	}
}

func BenchmarkList_IndexOf(b *testing.B) {
	list := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < b.N; i++ {
		_ = list.IndexOf(i)
	}
}

func TestList_Remove(t *testing.T) {
	intValues := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	list := New(intValues...)

	list.Remove(0)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, list.Values())

	list.Remove(3)
	assert.Equal(t, []int{1, 2, 3, 5, 6, 7, 8, 9}, list.Values())
}

func BenchmarkList_Remove(b *testing.B) {
	list := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < b.N; i++ {
		list.Remove(i)
	}
}

func TestList_Insert(t *testing.T) {
	list := New[int]()

	list.Insert(999, 0)
	assert.Equal(t, 0, list.size)

	list.Insert(0, 0)
	assert.Equal(t, 1, list.size)
	assert.Equal(t, []int{0}, list.Values())

	list.Insert(0, 1)
	assert.Equal(t, 2, list.size)
	assert.Equal(t, []int{1, 0}, list.Values())

	list.Insert(0, 3, 2)
	assert.Equal(t, 4, list.size)
	assert.Equal(t, []int{3, 2, 1, 0}, list.Values())
}

func BenchmarkList_Insert(b *testing.B) {
	list := New[int]()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		list.Insert(rand.Intn(b.N), i)
	}
}

func TestList_Set(t *testing.T) {
	list := New[int]()

	list.Set(999, 0)
	assert.Equal(t, 0, list.size)

	list.Set(0, 0)
	assert.Equal(t, 1, list.size)
	assert.Equal(t, []int{0}, list.Values())

	list.Set(0, 1)
	assert.Equal(t, 1, list.size)
	assert.Equal(t, []int{1}, list.Values())
}

func BenchmarkList_Set(b *testing.B) {
	list := New[int]()
	for i := 0; i < b.N; i++ {
		list.Set(0, i)
	}
}
