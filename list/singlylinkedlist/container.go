package singlylinkedlist

import (
	"fmt"
	"strings"

	"github.com/wwwangxc/container"
)

var _ container.Container[string] = (*List[string])(nil)

// IsEmpty will return true when container has no element
func (l *List[T]) IsEmpty() bool {
	return l.size == 0
}

// Size will return the number of elements in the container
func (l *List[T]) Size() int {
	return l.size
}

// Clear the elements inside the container
func (l *List[T]) Clear() {
	l.size = 0
	l.head = nil
	l.tail = nil
}

// Values will return the collection of all element values in the container
func (l *List[T]) Values() []T {
	values := make([]T, 0, l.size)
	l.Each(func(_ int, value T) bool {
		values = append(values, value)
		return true
	})

	return values
}

// String returns a string representation of singly linked list
func (l *List[T]) String() string {
	valueStr := make([]string, 0, l.size)
	l.Each(func(_ int, value T) bool {
		valueStr = append(valueStr, fmt.Sprintf("%v", value))
		return true
	})

	return fmt.Sprintf("SinglyLinkedList[%s]", strings.Join(valueStr, ", "))
}
