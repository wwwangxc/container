package minheap

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// Array implementation of heap
type MinHeap[T constraints.Ordered] struct {
	elements []T
	size     int
}

// Returns a new MinHeap implementation
func New[T constraints.Ordered](values ...T) *MinHeap[T] {
	minhp := &MinHeap[T]{
		elements: make([]T, 0),
		size:     0,
	}

	// Add elements to the heap and heapify
	for _, item := range values {
		minhp.elements = append(minhp.elements, item)
		minhp.size++
		minhp.heapifyUp()
	}
	return minhp
}

// add a value to the heap
func (mh *MinHeap[T]) add(value T) {
	mh.elements = append(mh.elements, value)
	mh.size++
	mh.heapifyUp()
}

// remove and return the top (minimum) value in the heap
func (mh *MinHeap[T]) poll() (T, error) {
	if mh.IsEmpty() {
		return *new(T), errors.New("heap is empty")
	}
	item := mh.elements[0]
	mh.elements[0] = mh.elements[mh.size-1]
	mh.size--
	mh.heapifyDown()
	return item, nil
}

// peek at the top node
func (mh *MinHeap[T]) peek() (T, error) {
	if mh.IsEmpty() {
		return *new(T), errors.New("heap is empty")
	}

	return mh.elements[0], nil
}

// swap one element with another
func (mh *MinHeap[T]) swap(idx1 int, idx2 int) {
	temp := mh.elements[idx1]
	mh.elements[idx1] = mh.elements[idx2]
	mh.elements[idx2] = temp
}

// Put the last element to its correct position
func (mh *MinHeap[T]) heapifyUp() {
	index := mh.size - 1

	for mh.HasParent(index) && mh.Parent(index) > mh.elements[index] {
		pIndex := mh.GetParentIndex(index)
		mh.swap(pIndex, index)
		index = pIndex
	}
}

// Put the first element to its correct position
func (mh *MinHeap[T]) heapifyDown() {
	index := 0

	for mh.HasLeftChild(index) {
		smChildIdx := mh.GetLeftChildIndex(index)
		if mh.HasRightChild(index) && mh.RightChild(index) < mh.LeftChild(index) {
			smChildIdx = mh.GetRightChildIndex(index)
		}

		if mh.elements[index] > mh.elements[smChildIdx] {
			mh.swap(index, smChildIdx)
		} else {
			break
		}
		index = smChildIdx
	}

}
