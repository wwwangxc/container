package heap

import (
	"github.com/wwwangxc/container"
	"github.com/wwwangxc/container/heap/minheap"
	"golang.org/x/exp/constraints"
)

// Return a new heap
func NewMinHeap[T constraints.Ordered](elements ...T) container.Heap[T] {
	return minheap.New(elements...)
}
