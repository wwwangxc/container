package hashset

var empty = struct{}{}

// Set implement of the set by a hash map
// Not thread safety
// Reference: https://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
type Set[T comparable] struct {
	m map[T]struct{}
}

// New return implement of the set by a hash map
func New[T comparable](elements ...T) *Set[T] {
	set := &Set[T]{}
	set.Add(elements...)

	return set
}
