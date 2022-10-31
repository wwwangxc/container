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

	"github.com/wwwangxc/container/list/arraylist"
)

func main() {
	list := arraylist.New("a", "b", "c") // ["a", "b", "c"]
	_, _ = list.Get(0)                   // "a", true
	_, _ = list.Get(3)                   // "", false
	_ = list.Contains("a")               // true
	_ = list.Contains("d")               // false
	_ = list.IndexOf("a")                // 0
	_ = list.IndexOf("d")                // -1
	list.Append("d", "e")                // ["a", "b", "c", "d", "e"]
	list.Remove(3)                       // ["a", "b", "c", "e"]
	list.Remove(3)                       // ["a", "b", "c"]
	list.Insert(3, "f", "g")             // ["a", "b", "c", "f", "g"]
	list.Remove(3)                       // ["a", "b", "c", "f"]
	list.Remove(3)                       // ["a", "b", "c"]
	list.Insert(0, "h", "i")             // ["h", "i", "a", "b", "c"]
	list.Set(0, "z")                     // ["z", "i", "a", "b", "c"]
	list.Set(5, "z")                     // ["z", "i", "a", "b", "c", "z"]
	list.SortBy(func(i, j string) bool { // ["a", "b", "c", "i", "z", "z"]
		return i < j
	})
	_ = list.IsEmpty() // false
	_ = list.Size()    // 6
	_ = list.Values()  // ["a", "b", "c", "i", "z", "z"]
	_ = list.String()  // ArrayList[a, b, c, i, z, z]
	list.Clear()       // []

	list = arraylist.New("a", "b", "c")

	list.Each(func(index int, value string) {
		fmt.Printf("Index: %d, Value: %s", index, value)
	})

	// return true
	_ = list.Any(func(index int, value string) bool {
		return value == "a"
	})

	// return false
	_ = list.All(func(index int, value string) bool {
		return value == "a"
	})

	// return 0, "a", true
	_, _, _ = list.Find(func(index int, value string) bool {
		return value == "a"
	})

	// return ["a"]
	_ = list.Select(func(index int, value string) bool {
		return value == "a"
	})
}

```