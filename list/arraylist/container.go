package arraylist

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
	l.elements = []T{}
}

// Values will return the collection of all element values in the container
func (l *List[T]) Values() []T {
	if l.size == 0 {
		return []T{}
	}

	elements := make([]T, l.size)
	copy(elements, l.elements[:l.size])
	return elements
}

// String returns a string representation of array list
func (l *List[T]) String() string {
	valueStr := make([]string, 0, l.size)
	for _, v := range l.Values() {
		valueStr = append(valueStr, fmt.Sprintf("%v", v))
	}

	return fmt.Sprintf("ArrayList[%s]", strings.Join(valueStr, ", "))
}
