package doublylinkedlist

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

	return l.Values()[index], true
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

// IndexOf the value
func (l *List[T]) IndexOf(value T) int {
	for i, v := range l.Values() {
		if v == value {
			return i
		}
	}

	return -1
}

// Remove the value
func (l *List[T]) Remove(index int) {
	if l.outOfRange(index) {
		return
	}

	if l.size == 1 {
		l.Clear()
		return
	}

	e := l.head
	for index > 0 {
		e = e.next
		index--
	}

	switch e {
	case l.head:
		l.head = e.next
		l.head.prev = nil
	case l.tail:
		l.tail = l.tail.prev
		l.tail.next = nil
	default:
		e.prev.next = e.next
		e.next.prev = e.prev
	}

	l.size--
}

// SortBy custom method
func (l *List[T]) SortBy(less func(i, j T) bool) {
	if l.size < 2 {
		return
	}

	values := l.Values()
	sort.Slice(values, func(i, j int) bool {
		return less(values[i], values[j])
	})

	l.Clear()
	l.Append(values...)
}

// Append value to list
func (l *List[T]) Append(values ...T) {
	for _, v := range values {
		l.append(v)
	}
}

// Insert values from index position
func (l *List[T]) Insert(index int, values ...T) {
	switch {
	case len(values) == 0:
		return
	case index == l.size:
		l.Append(values...)
		return
	case l.outOfRange(index):
		return
	}

	before := &element[T]{next: l.head}
	for index > 0 {
		before = before.next
		index--
	}

	head := &element[T]{value: values[0]}
	tail := head
	for i := 1; i < len(values); i++ {
		tail.next = &element[T]{value: values[i], prev: tail}
		tail = tail.next
	}

	e := before.next
	before.next = head
	head.prev = before
	tail.next = e
	e.prev = tail
	l.size += len(values)

	if e == l.head {
		l.head = head
		l.head.prev = nil
	}
}

// Set will override the value at the index position
func (l *List[T]) Set(index int, value T) {
	switch {
	case index == l.size:
		l.append(value)
		return
	case l.outOfRange(index):
		return
	}

	e := l.head
	for index > 0 {
		e = e.next
		index--
	}

	e.value = value
}
