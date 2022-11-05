# List

List is the base feature provided by all list structures, it stores values and the values can be repeated.

Implementes [Container](../README.md#container) interface.

Wiki: [https://en.wikipedia.org/wiki/List_%28abstract_data_type%29](https://en.wikipedia.org/wiki/List_%28abstract_data_type%29)

```go
// List is the base feature provided by all list structures
type List[T any] interface {

	// Get element value by index
	Get(index int) (T, bool)

	// Contains all the given values
	Contains(values ...T) bool

	// IndexOf the value
	IndexOf(value T) int

	// Remove element by index
	Remove(index int)

	// SortBy custom method
	SortBy(less func(i, j T) bool)

	// Append value to list
	Append(values ...T)

	// Insert values from index position
	Insert(index int, values ...T)

	// Set will override the value at the index position
	Set(index int, value T)

	// Container is the base feature provided by all data structures
	Container[T]

	// Enumerable provides enumerable functions for the containers
	Enumerable[int, T]
}

```

## Array List

Array is a data structure consisting of a collection of elements (values or variables), each identified by at least one array index.

Implements [Container](../README.md#container)、[Enumerable](../README.md#enumerable)、[List](#list) interface.

Wiki: [https://en.wikipedia.org/wiki/Array_(data_structure)](https://en.wikipedia.org/wiki/Array_(data_structure))

Usage by:

```go
package main

import (
	"fmt"

	"github.com/wwwangxc/container/list"
	"github.com/wwwangxc/container/list/arraylist"
)

func main() {
	array := list.NewArray("a", "b", "c") // ["a", "b", "c"]
	_, _ = array.Get(0)                   // "a", true
	_, _ = array.Get(3)                   // "", false
	_ = array.Contains("a")               // true
	_ = array.Contains("d")               // false
	_ = array.IndexOf("a")                // 0
	_ = array.IndexOf("d")                // -1
	array.Append("d", "e")                // ["a", "b", "c", "d", "e"]
	array.Remove(3)                       // ["a", "b", "c", "e"]
	array.Remove(3)                       // ["a", "b", "c"]
	array.Insert(3, "f", "g")             // ["a", "b", "c", "f", "g"]
	array.Remove(3)                       // ["a", "b", "c", "f"]
	array.Remove(3)                       // ["a", "b", "c"]
	array.Insert(0, "h", "i")             // ["h", "i", "a", "b", "c"]
	array.Set(0, "z")                     // ["z", "i", "a", "b", "c"]
	array.Set(5, "z")                     // ["z", "i", "a", "b", "c", "z"]
	array.SortBy(func(i, j string) bool { // ["a", "b", "c", "i", "z", "z"]
		return i < j
	})
	_ = array.IsEmpty() // false
	_ = array.Size()    // 6
	_ = array.Values()  // ["a", "b", "c", "i", "z", "z"]
	_ = array.String()  // ArrayList[a, b, c, i, z, z]
	array.Clear()       // []

	array = arraylist.New("a", "b", "c")

	array.Each(func(index int, value string) bool {
		fmt.Printf("Index: %d, Value: %s", index, value)

		// return true -> enter next loop
		// return false -> break loop
		return true
	})

	// return true
	_ = array.Any(func(index int, value string) bool {
		return value == "a"
	})

	// return false
	_ = array.All(func(index int, value string) bool {
		return value == "a"
	})

	// return 0, "a", true
	_, _, _ = array.Find(func(index int, value string) bool {
		return value == "a"
	})

	// return ["a"]
	_ = array.Select(func(index int, value string) bool {
		return value == "a"
	})
}
```

## Singly Linked List

Singly linked lists contain nodes which have a 'value' field as well as 'next' field, which points to the next node in line of nodes.
Operations that can be performed on singly linked lists include insertion, deletion and traversal.

<p align="center"><img src="https://upload.wikimedia.org/wikipedia/commons/6/6d/Singly-linked-list.svg"/></p>

Implements [Container](../README.md#container)、[Enumerable](../README.md#enumerable)、[List](#list) interface.

Wiki: [https://en.wikipedia.org/wiki/Linked_list#Singly_linked_list](https://en.wikipedia.org/wiki/Linked_list#Singly_linked_list)

Usage by:

```go
package main

import (
	"fmt"

	"github.com/wwwangxc/container/list"
	"github.com/wwwangxc/container/list/singlylinkedlist"
)

func main() {
	singly := list.NewSingly("a", "b", "c") // ["a", "b", "c"]
	_, _ = singly.Get(0)                    // "a", true
	_, _ = singly.Get(3)                    // "", false
	_ = singly.Contains("a")                // true
	_ = singly.Contains("d")                // false
	_ = singly.IndexOf("a")                 // 0
	_ = singly.IndexOf("d")                 // -1
	singly.Append("d", "e")                 // ["a", "b", "c", "d", "e"]
	singly.Remove(3)                        // ["a", "b", "c", "e"]
	singly.Remove(3)                        // ["a", "b", "c"]
	singly.Insert(3, "f", "g")              // ["a", "b", "c", "f", "g"]
	singly.Remove(3)                        // ["a", "b", "c", "f"]
	singly.Remove(3)                        // ["a", "b", "c"]
	singly.Insert(0, "h", "i")              // ["h", "i", "a", "b", "c"]
	singly.Set(0, "z")                      // ["z", "i", "a", "b", "c"]
	singly.Set(5, "z")                      // ["z", "i", "a", "b", "c", "z"]
	singly.SortBy(func(i, j string) bool {  // ["a", "b", "c", "i", "z", "z"]
		return i < j
	})
	_ = singly.IsEmpty() // false
	_ = singly.Size()    // 6
	_ = singly.Values()  // ["a", "b", "c", "i", "z", "z"]
	_ = singly.String()  // ArrayList[a, b, c, i, z, z]
	singly.Clear()       // []

	singly = singlylinkedlist.New("a", "b", "c")

	singly.Each(func(index int, value string) bool {
		fmt.Printf("Index: %d, Value: %s", index, value)

		// return true -> enter next loop
		// return false -> break loop
		return true
	})

	// return true
	_ = singly.Any(func(index int, value string) bool {
		return value == "a"
	})

	// return false
	_ = singly.All(func(index int, value string) bool {
		return value == "a"
	})

	// return 0, "a", true
	_, _, _ = singly.Find(func(index int, value string) bool {
		return value == "a"
	})

	// return ["a"]
	_ = singly.Select(func(index int, value string) bool {
		return value == "a"
	})
}
```

## Doubly Linked List

Doubly linked list is a linked data structure that consists of a set of sequentially linked records called nodes. 
Each node contains three fields: two link fields (references to the previous and to the next node in the sequence of nodes) and one data field. 

<p align="center"><img src="https://upload.wikimedia.org/wikipedia/commons/5/5e/Doubly-linked-list.svg"/></p>

Implements [Container](../README.md#container)、[Enumerable](../README.md#enumerable)、[List](#list) interface.

Wiki: [https://en.wikipedia.org/wiki/Doubly_linked_list](https://en.wikipedia.org/wiki/Doubly_linked_list)

Usage by:

```go
package main

import (
	"fmt"

	"github.com/wwwangxc/container/list"
	"github.com/wwwangxc/container/list/doublylinkedlist"
)

func main() {
	doubly := list.NewDoubly("a", "b", "c") // ["a", "b", "c"]
	_, _ = doubly.Get(0)                    // "a", true
	_, _ = doubly.Get(3)                    // "", false
	_ = doubly.Contains("a")                // true
	_ = doubly.Contains("d")                // false
	_ = doubly.IndexOf("a")                 // 0
	_ = doubly.IndexOf("d")                 // -1
	doubly.Append("d", "e")                 // ["a", "b", "c", "d", "e"]
	doubly.Remove(3)                        // ["a", "b", "c", "e"]
	doubly.Remove(3)                        // ["a", "b", "c"]
	doubly.Insert(3, "f", "g")              // ["a", "b", "c", "f", "g"]
	doubly.Remove(3)                        // ["a", "b", "c", "f"]
	doubly.Remove(3)                        // ["a", "b", "c"]
	doubly.Insert(0, "h", "i")              // ["h", "i", "a", "b", "c"]
	doubly.Set(0, "z")                      // ["z", "i", "a", "b", "c"]
	doubly.Set(5, "z")                      // ["z", "i", "a", "b", "c", "z"]
	doubly.SortBy(func(i, j string) bool {  // ["a", "b", "c", "i", "z", "z"]
		return i < j
	})
	_ = doubly.IsEmpty() // false
	_ = doubly.Size()    // 6
	_ = doubly.Values()  // ["a", "b", "c", "i", "z", "z"]
	_ = doubly.String()  // ArrayList[a, b, c, i, z, z]
	doubly.Clear()       // []

	doubly = doublylinkedlist.New("a", "b", "c")

	doubly.Each(func(index int, value string) bool {
		fmt.Printf("Index: %d, Value: %s", index, value)

		// return true -> enter next loop
		// return false -> break loop
		return true
	})

	// return true
	_ = doubly.Any(func(index int, value string) bool {
		return value == "a"
	})

	// return false
	_ = doubly.All(func(index int, value string) bool {
		return value == "a"
	})

	// return 0, "a", true
	_, _, _ = doubly.Find(func(index int, value string) bool {
		return value == "a"
	})

	// return ["a"]
	_ = doubly.Select(func(index int, value string) bool {
		return value == "a"
	})
}
```
