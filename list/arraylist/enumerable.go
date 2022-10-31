package arraylist

import "github.com/wwwangxc/container"

var _ container.Enumerable[int, string] = (*List[string])(nil)

// Each calls the given function once for each element and passing that
// element's index and value
func (l *List[T]) Each(f func(index int, value T)) {
	for i, v := range l.elements {
		f(i, v)
	}
}

// Any calls the given function once for each element
//
// Return true if the given function returns true once
// Return false if the given function always returns false for all elements
func (l *List[T]) Any(f func(index int, value T) bool) bool {
	for i, v := range l.elements {
		if f(i, v) {
			return true
		}
	}

	return false
}

// All calls the given function once for each element
//
// Return true if the given function always returns true for all elements
// Return false if the given function returns false once
func (l *List[T]) All(f func(index int, value T) bool) bool {
	for i, v := range l.elements {
		if !f(i, v) {
			return false
		}
	}

	return true
}

// Find calls the given function once for each element
//
// Return [index, value, true] when the given function return true
// for the first time
func (l *List[T]) Find(f func(index int, value T) bool) (int, T, bool) {
	for i, v := range l.elements {
		if f(i, v) {
			return i, v, true
		}
	}

	var value T
	return -1, value, false
}

// Select calls the given function once for each element
//
// It will return all elements for which the given function returns true
func (l *List[T]) Select(f func(index int, value T) bool) []T {
	var values []T
	for i, v := range l.elements {
		if f(i, v) {
			values = append(values, v)
		}
	}

	return values
}
