package container

import "fmt"

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
