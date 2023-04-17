package heap

import (
	"github.com/wwwangxc/container"
	"github.com/wwwangxc/container/heap/maxheap"
	"github.com/wwwangxc/container/heap/minheap"
	"golang.org/x/exp/constraints"
)

// Return a new minheap
func NewMinHeap[T constraints.Ordered](elements ...T) container.Heap[T] {
	return minheap.New(elements...)
}

// Return a nee maxheap
func NewMaxHeap[T constraints.Ordered](elements ...T) container.Heap[T] {
    return maxheap.New(elements...)
}
