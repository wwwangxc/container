package maxheap

import (
	"github.com/wwwangxc/container"
)

var _ container.Heap[int] = (*MaxHeap[int])(nil)

// Add and element to the Heap
func (mh *MaxHeap[T]) Add(element T) {
    mh.add(element)
}

// Peek at the top element of the Heap
func (mh *MaxHeap[T]) Peek() (T, error) {
    return mh.peek()
}

// Pop out the top element of the Heap
func (mh *MaxHeap[T]) Poll() (T, error) {
    return mh.poll()
}

// Get the parent index of a node
func (mh *MaxHeap[T]) GetParentIndex(child int) int {
    return (child - 1) / 2
}

// Get the left child index of a given node
func (mh *MaxHeap[T]) GetLeftChildIndex(parent int) int {
    return 2 * parent + 1
}

// Get the right child index of a given node
func (mh *MaxHeap[T]) GetRightChildIndex(parent int) int {
    return 2 * parent + 2
}

// Get the left child of the given node
func (mh *MaxHeap[T]) LeftChild(parent int) T {
    return mh.elements[mh.GetLeftChildIndex(parent)]
}

// Get the right child of the given node
func (mh *MaxHeap[T]) RightChild(parent int) T {
    return mh.elements[mh.GetRightChildIndex(parent)]
}

// Get the parent of the given node
func (mh *MaxHeap[T]) Parent(child int) T {
    return mh.elements[mh.GetParentIndex(child)]
}

// Check if current node has a right child
func (mh *MaxHeap[T]) HasRightChild(parent int) bool {
    return mh.GetRightChildIndex(parent) < mh.size
}

// Check if the current node has a left child
func (mh *MaxHeap[T]) HasLeftChild(parent int) bool {
    return mh.GetLeftChildIndex(parent) < mh.size
}

// Check if the current node has a parent
func (mh *MaxHeap[T]) HasParent(child int) bool {
    return mh.GetParentIndex(child) >= 0
}
