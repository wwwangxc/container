package maxheap

import (
	"fmt"
	"strings"
)

// Return true when there are no nodes in the heap
func (mh *MaxHeap[T]) IsEmpty() bool {
    return mh.size == 0
}

// Returns the size of the heap
func (mh *MaxHeap[T]) Size() int {
    return mh.size
}

// Remove all elements from the heap
func (mh *MaxHeap[T]) Clear() {
    mh.size = 0;
    mh.elements = []T{}
}

// Return all the elements of the heap
func (mh *MaxHeap[T]) Values() []T {
    if mh.size == 0 {
		return []T{}
	}

	elements := make([]T, mh.size)
	copy(elements, mh.elements[:mh.size])
	return elements
}

// Returns a string representation of the heap
// Improve to show a tree like structure?
//     p
//   /  \
//  cl   cr
func (mh *MaxHeap[T]) String() string {
    valueStr := make([]string, 0, mh.size)
	for _, v := range mh.Values() {
		valueStr = append(valueStr, fmt.Sprintf("%v", v))
	}

	return fmt.Sprintf("MaxHeap [%s]", strings.Join(valueStr, ", "))
}


