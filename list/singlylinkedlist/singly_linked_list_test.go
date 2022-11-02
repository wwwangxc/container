package singlylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	intValues := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := New(intValues...)

	assert.NotNil(t, got)
	assert.Equal(t, got.size, len(intValues))
	assert.Equal(t, intValues, got.Values())
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
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6}, list.Values())
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
