package container

// Enumerable provides enumerable functions for the containers
type Enumerable[T comparable, U any] interface {

	// Each calls the given function once for each element and passing that
	// element's index(or key) and value
	//
	// Enter the next loop when it returns true
	// Break loop when it returns false
	Each(func(T, U) bool)

	// Any calls the given function once for each element
	//
	// Return true if the given function returns true once
	// Return false if the given function always returns false for all elements
	Any(func(T, U) bool) bool

	// All calls the given function once for each element
	//
	// Return true if the given function always returns true for all elements
	// Return false if the given function returns false once
	All(func(T, U) bool) bool

	// First calls the given function once for each element
	//
	// Return [index(or key), value, true] when the given function return true
	// for the first time
	Find(func(T, U) bool) (T, U, bool)

	// Select calls the given function once for each element
	//
	// It will return all elements for which the given function returns true
	Select(func(T, U) bool) []U
}
