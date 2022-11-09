# Set

`Set` is the base feature provided by all set structures.

A set is an abstract data type that can store unique values, without any particular order. 
It is a computer implementation of the mathematical concept of a finite set.

Wiki: [https://en.wikipedia.org/wiki/Set_%28abstract_data_type%29](https://en.wikipedia.org/wiki/Set_%28abstract_data_type%29)

```go
// Set is the base feature provided by all set structures
type Set[T comparable] interface {

	// Add all given elements into the set
	Add(elements ...T)

	// Remove all given elements from the set
	Remove(elements ...T)

	// Contains all given elements
	Contains(elements ...T) bool

	// Union returns the union of two sets.
	// All elements of the new set exist in the current set or exist in the target set.
	// Reference: https://en.wikipedia.org/wiki/Union_(set_theory)
	Union(target Set[T]) Set[T]

	// Intersection returns the intersection between two sets.
	// All elements of the new set exist in both current set and target set.
	// Reference: https://en.wikipedia.org/wiki/Intersection_(set_theory)
	Intersection(target Set[T]) Set[T]

	// Difference returns the difference between two sets.
	// All elements of the new set exist in the current set but not in the target set.
	// Reference: https://proofwiki.org/wiki/Definition:Set_Difference
	Difference(target Set[T]) Set[T]

	// Container is the base feature provided by all data structures
	Container[T]
}
```

## Hash Set

`HashSet` is an unordered collection containing unique elements. 

It has the standard collection operations `Add`, `Remove`, `Contains`, but since it uses a hash-based implementation, these operations are O(1).

`HashSet` also provides standard set operations such as union, intersection, and symmetric difference.

Implements [Container](../README.md#container)、[Set](#set) interface.

Usage by :

```go
package main

import (
	"github.com/wwwangxc/container/set"
	"github.com/wwwangxc/container/set/hashset"
)

func main() {
	set := set.NewHashSet(1, 2, 3) // {1, 2, 3}
	set.Add(4)                     // {1, 2, 3, 4}
	set.Remove(4)                  // {1, 2, 3}
	_ = set.Contains(1)            // true
	_ = set.Contains(1, 2, 3)      // true
	_ = set.Contains(4)            // false

	target := hashset.New(4, 5, 6) // {4, 5, 6}
	_ = set.Union(target)          // {1, 2, 3, 4, 5, 6}

	target = hashset.New(3, 4, 5) // {3, 4, 5}
	_ = set.Intersection(target)  // {3}

	target = hashset.New(3, 4, 5) // {3, 4, 5}
	_ = set.Difference(target)    // {1, 2}

	_ = set.String()  // HashSet[1, 2, 3]
	_ = set.IsEmpty() // false
	_ = set.Size()    // 3
	_ = set.Values()  // [1, 2, 3]
	set.Clear()       // {}
}
```
