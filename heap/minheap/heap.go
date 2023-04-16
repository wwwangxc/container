package minheap

import (
	"github.com/wwwangxc/container"
)

var _ container.Heap[int] = (*MinHeap[int])(nil)

func (mh *MinHeap[T]) Add(element T) {
    mh.add(element)
}

func (mh *MinHeap[T]) Peek() (T, error) {
    return mh.peek()
}

func (mh *MinHeap[T]) Poll() (T, error) {
    return mh.poll()
}

func (mh *MinHeap[T]) GetParentIndex(child int) int {
    return (child - 1) / 2
}

func (mh *MinHeap[T]) GetLeftChildIndex(parent int) int {
    return 2 * parent + 1
}

func (mh *MinHeap[T]) GetRightChildIndex(parent int) int {
    return 2 * parent + 2
}

func (mh *MinHeap[T]) LeftChild(parent int) T {
    return mh.elements[mh.GetLeftChildIndex(parent)]
}

func (mh *MinHeap[T]) RightChild(parent int) T {
    return mh.elements[mh.GetRightChildIndex(parent)]
}

func (mh *MinHeap[T]) Parent(child int) T {
    return mh.elements[mh.GetParentIndex(child)]
}

func (mh *MinHeap[T]) HasRightChild(parent int) bool {
    return mh.GetRightChildIndex(parent) < mh.size
}

func (mh *MinHeap[T]) HasLeftChild(parent int) bool {
    return mh.GetLeftChildIndex(parent) < mh.size
}

func (mh *MinHeap[T]) HasParent(child int) bool {
    return mh.GetParentIndex(child) >= 0
}


