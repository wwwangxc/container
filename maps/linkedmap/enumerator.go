package linkedmap

import "github.com/wwwangxc/container"

var _ container.Enumerator[string, string] = (*Map[string, string])(nil)

// Each calls the given function once for each element and passing that
// element's key and value
//
// Enter the next loop when it returns true.
// Break loop when it returns false.
func (m *Map[K, V]) Each(f func(k K, v V) bool) {
	m.list.Each(func(_ int, k K) bool {
		return f(k, m.m[k])
	})
}

// Any calls the given function once for each element
//
// Return true if the given function returns true once.
// Return false if the given function always returns false for all elements.
func (m *Map[K, V]) Any(f func(k K, v V) bool) bool {
	ok := true
	m.list.Each(func(_ int, k K) bool {
		if ok = f(k, m.m[k]); ok {
			// break loop
			return false
		}

		// enter the next loop
		return true
	})

	return ok
}

// All calls the given function once for each element
//
// Return true if the given function always returns true for all elements>
// Return false if the given function returns false once.
func (m *Map[K, V]) All(f func(k K, v V) bool) bool {
	ok := true
	m.list.Each(func(_ int, k K) bool {
		if ok = ok && f(k, m.m[k]); !ok {
			// break loop
			return false
		}

		// enter the next loop
		return true
	})

	return ok
}

// Find calls the given function once for each element
//
// Return [key, value, true] when the given function return true
// for the first time
func (m *Map[K, V]) Find(f func(k K, v V) bool) (K, V, bool) {
	var key K
	var value V
	exist := false

	m.list.Each(func(_ int, k K) bool {
		if ok := f(k, m.m[k]); ok {
			key = k
			value = m.m[k]
			exist = true
			// break loop
			return false
		}

		// enter the next loop
		return true
	})

	return key, value, exist
}

// Select calls the given function once for each element
//
// It will return all elements for which the given function returns true
func (m *Map[K, V]) Select(f func(k K, v V) bool) []V {
	var values []V
	m.list.Each(func(_ int, k K) bool {
		if f(k, m.m[k]) {
			values = append(values, m.m[k])
		}

		// enter the next loop
		return true
	})

	return values
}
