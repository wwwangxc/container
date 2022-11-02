package arraylist

import (
	"sort"

	"github.com/wwwangxc/container"
)

var _ container.List[string] = (*List[string])(nil)

// Get element value by index
func (l *List[T]) Get(index int) (T, bool) {
	if l.outOfRange(index) {
		var zero T
		return zero, false
	}

	return l.elements[index], true
}

// Contains all the given values
func (l *List[T]) Contains(values ...T) bool {
	for _, v := range values {
		if l.IndexOf(v) < 0 {
			return false
		}
	}

	return true
}

// IndexOf the value in current list
func (l *List[T]) IndexOf(value T) int {
	for i, v := range l.Values() {
		if v == value {
			return i
		}
	}

	return -1
}

// Remove element by index
// Current operation will trigger shrinking when the number of elements is less than the threshold
func (l *List[T]) Remove(index int) {
	if l.outOfRange(index) {
		return
	}

	copy(l.elements[index:], l.elements[index+1:l.size])
	l.size--
	l.tryShrink()
}

// SortBy custom method
func (l *List[T]) SortBy(less func(i, j T) bool) {
	if less == nil {
		return
	}

	sort.Slice(l.elements, func(i, j int) bool {
		return less(l.elements[i], l.elements[j])
	})
}

// Append value to list
func (l *List[T]) Append(values ...T) {
	for _, v := range values {
		l.append(v)
	}
}

// Insert values from index position
func (l *List[T]) Insert(index int, values ...T) {
	if len(values) == 0 {
		return
	}

	if l.outOfRange(index) {
		if index == l.size {
			l.Append(values...)
		}
		return
	}

	left := l.elements[:index]
	right := l.elements[index:l.size]
	elements := make([]T, 0, len(left)+len(values)+len(right))
	elements = append(elements, left...)
	elements = append(elements, values...)
	elements = append(elements, right...)
	l.elements = elements
	l.size += len(values)
}

// Set will override the value at the index position
func (l *List[T]) Set(index int, value T) {
	if l.outOfRange(index) {
		if index == l.size {
			l.append(value)
		}
		return
	}

	l.elements[index] = value
}
