package arraylist

import (
	"sort"

	"github.com/wwwangxc/container"
)

var _ container.List[string] = (*List[string])(nil)

// Get element value by index
func (a *List[T]) Get(index int) (T, bool) {
	if a.outOfRange(index) {
		var zero T
		return zero, false
	}

	return a.elements[index], true
}

// Contains all the given values
func (a *List[T]) Contains(values ...T) bool {
	for _, v := range values {
		if a.IndexOf(v) < 0 {
			return false
		}
	}

	return true
}

// IndexOf the value in current list
func (a *List[T]) IndexOf(value T) int {
	for i, v := range a.Values() {
		if v == value {
			return i
		}
	}

	return -1
}

// Remove element by index
// Current operation will trigger shrinking when the number of elements is less than the threshold
func (a *List[T]) Remove(index int) {
	if a.outOfRange(index) {
		return
	}

	copy(a.elements[index:], a.elements[index+1:a.size])
	a.size--
	a.tryShrink()
}

// SortBy custom method
func (a *List[T]) SortBy(less func(i, j T) bool) {
	if less == nil {
		return
	}

	sort.Slice(a.elements, func(i, j int) bool {
		return less(a.elements[i], a.elements[j])
	})
}

// Append value to list
func (a *List[T]) Append(values ...T) {
	for _, v := range values {
		a.append(v)
	}
}

// Insert values from index position
func (a *List[T]) Insert(index int, values ...T) {
	if len(values) == 0 {
		return
	}

	if a.outOfRange(index) {
		if index == a.size {
			a.Append(values...)
		}
		return
	}

	left := a.elements[:index]
	right := a.elements[index:a.size]
	elements := make([]T, 0, len(left)+len(values)+len(right))
	elements = append(elements, left...)
	elements = append(elements, values...)
	elements = append(elements, right...)
	a.elements = elements
	a.size += len(values)
}

// Set will override the value at the index position
func (a *List[T]) Set(index int, value T) {
	if a.outOfRange(index) {
		if index == a.size {
			a.append(value)
		}
		return
	}

	a.elements[index] = value
}
