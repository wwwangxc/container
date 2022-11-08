package container

// List is the base feature provided by all list structures
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
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

	// Enumerator provides enumerator functions for the containers
	Enumerator[int, T]
}
