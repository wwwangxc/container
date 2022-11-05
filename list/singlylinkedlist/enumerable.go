package singlylinkedlist

import (
	"github.com/wwwangxc/container"
)

var _ container.Enumerable[int, string] = (*List[string])(nil)

// Each calls the given function once for each element and passing that
// element's index and value
//
// Enter the next loop when it returns true
// Break loop when it returns false
func (l *List[T]) Each(f func(index int, value T) bool) {
	index := -1
	for e := l.head; e != nil; e = e.next {
		index++
		if next := f(index, e.value); !next {
			break
		}
	}
}

// Any calls the given function once for each element
//
// Return true if the given function returns true once
// Return false if the given function always returns false for all elements
func (l *List[T]) Any(f func(index int, value T) bool) bool {
	ok := false
	l.Each(func(i int, v T) bool {
		if ok = f(i, v); ok {
			// break loop
			return false
		}

		// enter the next loo;
		return true
	})

	return ok
}

// All calls the given function once for each element
//
// Return true if the given function always returns true for all elements
// Return false if the given function returns false once
func (l *List[T]) All(f func(index int, value T) bool) bool {
	ok := true
	l.Each(func(i int, v T) bool {
		if ok = ok && f(i, v); !ok {
			// break loop
			return false
		}

		// enter the next loop
		return true
	})

	return ok
}

// First calls the given function once for each element
//
// Return [index(or key), value, true] when the given function return true
// for the first time
func (l *List[T]) Find(f func(index int, value T) bool) (int, T, bool) {
	index := -1
	var value T
	exist := false

	l.Each(func(i int, v T) bool {
		if ok := f(i, v); ok {
			index = i
			value = v
			exist = true

			// break loop
			return false
		}

		// enter the next loop
		return true
	})

	return index, value, exist
}

// Select calls the given function once for each element
//
// It will return all elements for which the given function returns true
func (l *List[T]) Select(f func(index int, value T) bool) []T {
	var values []T
	l.Each(func(i int, v T) bool {
		if f(i, v) {
			values = append(values, v)
		}

		// enter the next loop
		return true
	})

	return values
}
