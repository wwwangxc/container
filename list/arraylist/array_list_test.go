package arraylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	intValues := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := New(intValues...)

	assert.NotNil(t, got)
	assert.Equal(t, got.size, len(intValues))
	assert.Equal(t, cap(got.elements), len(intValues))
	assert.Equal(t, got.elements, intValues)
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}

func TestList_append(t *testing.T) {
	list := New(0, 1, 2, 3, 4, 5)
	list.append(6)
	assert.Equal(t, list.size, 7)
	assert.Equal(t, list.elements, []int{0, 1, 2, 3, 4, 5, 6})

	list = &List[int]{
		elements: []int{0, 1, 2, 3, 4, 5},
		size:     0,
	}
	list.append(6)
	assert.Equal(t, list.size, 1)
	assert.Equal(t, list.elements, []int{6, 1, 2, 3, 4, 5})
}

func BenchmarkList_append(b *testing.B) {
	list := New[int]()
	for i := 0; i < b.N; i++ {
		list.append(i)
	}
}

func BenchmarkList_outOfRange(b *testing.B) {
	list := New[int]()
	for i := 0; i < b.N; i++ {
		list.append(i)
	}
}

func TestList_tryShrink(t *testing.T) {
	list := &List[int]{
		elements: make([]int, 0, 100),
		size:     0,
	}
	list.tryShrink()
	assert.Equal(t, int(float32(100)*thresholdShrink), cap(list.elements))

	list = &List[int]{
		elements: make([]int, 100),
		size:     100,
	}
	list.tryShrink()
	assert.Equal(t, cap(list.elements), 100)
}

func BenchmarkList_tryShrink(b *testing.B) {
	list := &List[int]{
		elements: make([]int, 0, 100),
		size:     0,
	}

	for i := 0; i < b.N; i++ {
		list.tryShrink()
	}
}

func TestList_resize(t *testing.T) {
	list := &List[int]{
		elements: make([]int, 0, 100),
		size:     0,
	}
	list.resize(10)
	assert.Equal(t, len(list.elements), 0)
	assert.Equal(t, cap(list.elements), 10)

	list = &List[int]{
		elements: make([]int, 10, 100),
		size:     10,
	}
	list.resize(10)
	assert.Equal(t, len(list.elements), 10)
	assert.Equal(t, cap(list.elements), 10)
}

func BenchmarkList_resize(b *testing.B) {
	list := &List[int]{
		elements: make([]int, 0, 100),
		size:     0,
	}

	for i := 0; i < b.N; i++ {
		list.resize(25)
	}
}
