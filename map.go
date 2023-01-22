package container

// Map is the base feature provided by all map structures
//
// In computer science, an associative array, map, symbol table, or dictionary
// is an abstract data type composed of a collection of (key, value) pairs, such
// that each possible key appears just once in the collection.
//
// Reference: https://en.wikipedia.org/wiki/Map_(higher-order_function)
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
