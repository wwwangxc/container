package maxheap

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// Array implementation of heap
type MaxHeap[T constraints.Ordered] struct {
    elements []T
    size int
}

// Returns a new MinHeap implementation
func New[T constraints.Ordered](values ...T) *MaxHeap[T] {
    mxhp := &MaxHeap[T]{
        elements: make([]T, 0),
        size: 0,
    }

    // Add elements to the heap and heapify
    for _, item := range values {
        mxhp.elements = append(mxhp.elements, item)
        mxhp.size++
        mxhp.heapifyUp()
    }
    return mxhp
}

// add a value to the heap
func (mh *MaxHeap[T]) add(value T) {
    mh.elements = append(mh.elements, value)
    mh.size++
    mh.heapifyUp()
}

// remove and return the top (maximum) value in the heap
func (mh *MaxHeap[T]) poll() (T, error) {
    if mh.IsEmpty() {
        return *new(T), errors.New("heap is empty")
    }
    item := mh.elements[0]
    mh.elements[0] = mh.elements[mh.size - 1]
    mh.size--
    mh.heapifyDown()
    return item, nil
}

// peek at the top node
func (mh *MaxHeap[T]) peek() (T, error) {
    if mh.IsEmpty() {
        return *new(T), errors.New("heap is empty")
    }

    return mh.elements[0], nil
}

// swap one element with another
func (mh *MaxHeap[T]) swap(idx1 int, idx2 int) {
    temp := mh.elements[idx1]
    mh.elements[idx1] = mh.elements[idx2]
    mh.elements[idx2] = temp
}

// Put the last element to its correct position
func (mh *MaxHeap[T]) heapifyUp() {
    index := mh.size - 1

    for mh.HasParent(index) && mh.Parent(index) < mh.elements[index] {
        pIndex := mh.GetParentIndex(index)
        mh.swap(pIndex, index)
        index = pIndex
    }
}

// Put the first element to its correct position
func (mh *MaxHeap[T]) heapifyDown() {
    index := 0

    for mh.HasLeftChild(index) {
        gtChildIdx := mh.GetLeftChildIndex(index)
        if mh.HasRightChild(index) && mh.RightChild(index) > mh.LeftChild(index) {
            gtChildIdx = mh.GetRightChildIndex(index)
        }

        if mh.elements[index] < mh.elements[gtChildIdx] {
            mh.swap(index, gtChildIdx)
        } else {
            break
        }
        index = gtChildIdx
    }
}


