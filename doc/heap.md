# Heap

Heap is the base feature provided by all heap structures

In computer science, a heap is a specialized tree-based data structure which is essentially an almost complete binary tree that satisfies the heap property. in a max heap, for any given node C, if P is a parent node of C, then the key (the value) of P is greater than or equal to the key of C. In a min heap, the key of P is less than or equal to the key of C.
The node at the "top" of the heap (with no parents) is called the root node.

[Reference](https://en.wikipedia.org/wiki/Heap_(data_structure))

```go
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
```

## Min Heap

A min-heap is a data structure such that. - the data contained in each node is less than (or equal to) the data in that node's children.
[Reference](https://www.cs.cmu.edu/~tcortina/15-121sp10/Unit06B.pdf)

It supports `add`, `poll`, `peek` operations

```go
package main

import "github.com/wwwangxc/container/heap/minheap"

func main() {
    mh := minheap.New(1, 2, 3, 4)
    min, err := mh.Poll()   // returns (1, nil) and removes 1 from the heap
    mh.Add(7)   // adds 7 to the heap
    min, err = mh.Peek()    // returns (2, nil)

    fmt.Println(mh.IsEmpty())   // prints false

    fmt.Println(mh.String())    // prints "Minheap [2, 3, 4, 7]

    fmt.Println(mh.Size())  // prints 4

    items := mh.Values()    // returns the items of the heap as a slice

    mh.Clear()  // empties the heap
}
```

# Max Heap

A max-heap is a data structure Where the value of the root node is greater than or equal to either of its children.
[Reference](https://www.tutorialspoint.com/data_structures_algorithms/heap_data_structure.htm)

It supports `add`, `poll`, `peek` operations

```go
package main

import "github.com/wwwangxc/container/heap/maxheap"

func main() {
    mh := minheap.New(1, 2, 3, 4)
    min, err := mh.Poll()   // returns (4, nil) and removes 4 from the heap
    mh.Add(7)   // adds 7 to the heap
    min, err = mh.Peek()    // returns (7, nil)

    fmt.Println(mh.IsEmpty())   // prints false

    fmt.Println(mh.String())    // prints "Minheap [1, 2, 3, 7]

    fmt.Println(mh.Size())  // prints 4

    items := mh.Values()    // returns the items of the heap as a slice

    mh.Clear()  // empties the heap
}
```
