[![OSCS Status](https://www.oscs1024.com/platform/badge/wwwangxc/container.svg?size=small)](https://www.murphysec.com/dr/aDgidkK7d72WfN9WRi)
[![Go Report Card](https://goreportcard.com/badge/github.com/wwwangxc/container)](https://goreportcard.com/report/github.com/wwwangxc/container)
[![GoDoc](https://pkg.go.dev/badge/github.com/wwwangxc/container?status.svg)](https://pkg.go.dev/github.com/wwwangxc/container)

Various data structures implemented using Go generics. 🤗

## Support

- [List](doc/list.md#list)
  - [Array List](doc/list.md#array-list)

## Global Ability

### Container

It is the base feature provided by all data structures.

```go
// Container is the base feature provided by all data structures
type Container[T any] interface {
	fmt.Stringer

	// IsEmpty will return true when container has no element
	IsEmpty() bool

	// Size will return the number of elements in the container
	Size() int

	// Clear the elements inside the container
	Clear()

	// Values will return the collection of all element values in the container
	Values() []T
}
```

### Enumerable

It provides enumerable functions for the containers.

```go
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
```
