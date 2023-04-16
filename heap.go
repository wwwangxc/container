package container

import (
	"golang.org/x/exp/constraints"
)

// Heap is the base feature provided bu all heap structures
//
// In computer science, a heap is a specialized tree-based data structure
// which is essentially an almost complete binary tree that
// satisfies the heap property. in a max heap, for any given node C,
// if P is a parent node of C, then the key (the value) of
// P is greater than or equal to the key of C. In a min heap, the key of P is less than or equal to the key of C.
// The node at the "top" of the heap (with no parents) is called the root node.
//
// Reference: https://en.wikipedia.org/wiki/Heap_(data_structure)
type Heap[T constraints.Ordered] interface {

    // Add an elemnt the to Heap
    Add(element T)

    // Get the top Element of the heap without removing it
    Peek() (T, error)

    // Pop out the top element of the heap
    Poll() (T, error)

    // Get the index of the parent of given node
    GetParentIndex(child int) int

    // Get the index of the left child of the given node
    GetLeftChildIndex(parent int) int

    // Get the index of the right child of the given node
    GetRightChildIndex(parent int) int

    // Get the left child of a node
    LeftChild(parent int) T

    // Get the right child of a node
    RightChild(parent int) T

    // Get the parent node of a node
    Parent(child int) T

    // Check if the given node has a left child
    HasLeftChild(parent int) bool

    // Check if the given node has a right child
    HasRightChild(parent int) bool

    // Check if the given node has a parent
    HasParent(child int) bool

    // Container is the base feature provided by all data structures
    Container[T]
}
