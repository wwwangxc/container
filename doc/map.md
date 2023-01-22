# Map

`Map` is the base feature provided by all map structures.

In computer science, an associative array, map, symbol table, or dictionary
is an abstract data type composed of a collection of (key, value) pairs, such
that each possible key appears just once in the collection.

Wiki: [https://en.wikipedia.org/wiki/Map_(higher-order_function)](https://en.wikipedia.org/wiki/Map_(higher-order_function))

```go
// Map is the base feature provided by all map structures
type Map[K comparable, V any] interface {

	// Put the key and the value
	Put(k K, v V)

	// Get value by key
	Get(k K) (V, bool)

	// Delete value by ke
	Delete(k K)

	// Contains all given keys
	Contains(keys ...K) bool

	// Keys return all keys in the map
	Keys() []K

	// Container is the base feature provided by all data structures
	Container[V]

	// Enumerator provides enumerator functions for the containers
	Enumerator[K, V]
}
```

## Linked Map

In computer science, an associative array, map, symbol table, or dictionary
is an abstract data type that stores a collection of (key, value) pairs, 
such that each possible key appears at most once in the collection.

In mathematical terms an associative array is a function with finite domain.

It supports `lookup`, `remove`, and `insert` operations.

Implements [Container](../README.md#container)、[Enumerator](../README.md#enumerator)、[Map](#map) interface.

Wiki: [https://en.wikipedia.org/wiki/Associative_array](https://en.wikipedia.org/wiki/Associative_array)

Usage by:

```go
package main

import "github.com/wwwangxc/container/maps/linkedmap"

func main() {
	m := linkedmap.New[string, string]()
	m.Put("k_1", "v1")    // [k_1=>v_1]
	m.Put("k_2", "v2")    // [k_1=>v_1, k_2=>v_2]
	_, _ = m.Get("k_1")   // "v_1" true
	_ = m.Contains("k_1") // true
	_ = m.Keys()          // [k_1 k_2]
	_ = m.Values()        // [v_1 v_2]
	_ = m.IsEmpty()       // false
	_ = m.Size()          // 2
	_ = m.String()        // LinkedMap[k_1=>v_1, k_2=>v_2]

	m.Delete("key")       // LinkedMap[k_1=>v_1, k_2=>v_2]
	m.Delete("k_1")       // LinkedMap[k_2=>v_2]
	_, _ = m.Get("k_1")   // "" false
	_ = m.Contains("k_1") // false
	_ = m.Keys()          // [k_2]
	_ = m.Values()        // [v_2]
	_ = m.IsEmpty()       // false
	_ = m.Size()          // 1
	_ = m.String()        // LinkedMap[k_2=>v_2]

	m.Clear()             // []
	_, _ = m.Get("k_1")   // "" false
	_ = m.Contains("k_1") // false
	_ = m.Keys()          // []
	_ = m.Values()        // []
	_ = m.IsEmpty()       // true
	_ = m.Size()          // 0
	_ = m.String()        // LinkedMap[]
}
```
