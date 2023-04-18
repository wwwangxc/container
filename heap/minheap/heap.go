package minheap

import (
	"github.com/wwwangxc/container"
)

var _ container.Heap[int] = (*MinHeap[int])(nil)

// Add an element to the Heap
func (mh *MinHeap[T]) Add(element T) {
    mh.add(element)
}

// Peek at the top element of the Heap
func (mh *MinHeap[T]) Peek() (T, error) {
    return mh.peek()
}

// Pop out the top element of the Heap
func (mh *MinHeap[T]) Poll() (T, error) {
    return mh.poll()
}

// Get parent index of the given element
func (mh *MinHeap[T]) GetParentIndex(child int) int {
    return (child - 1) / 2
}

// Get left child index of the given element
func (mh *MinHeap[T]) GetLeftChildIndex(parent int) int {
    return 2 * parent + 1
}

// Get right child index of the given element
func (mh *MinHeap[T]) GetRightChildIndex(parent int) int {
    return 2 * parent + 2
}

// Get the left child of a given element
func (mh *MinHeap[T]) LeftChild(parent int) T {
    return mh.elements[mh.GetLeftChildIndex(parent)]
}

// Get the right child of the given element
func (mh *MinHeap[T]) RightChild(parent int) T {
    return mh.elements[mh.GetRightChildIndex(parent)]
}

// Get the parent of the given element
func (mh *MinHeap[T]) Parent(child int) T {
    return mh.elements[mh.GetParentIndex(child)]
}

// Check if current element has a right child
func (mh *MinHeap[T]) HasRightChild(parent int) bool {
    return mh.GetRightChildIndex(parent) < mh.size
}

// Check if current element has a left child
func (mh *MinHeap[T]) HasLeftChild(parent int) bool {
    return mh.GetLeftChildIndex(parent) < mh.size
}

// Check if current element has a parent
func (mh *MinHeap[T]) HasParent(child int) bool {
    return mh.GetParentIndex(child) >= 0
}
