package arraylist

import (
	"fmt"
	"strings"

	"github.com/wwwangxc/container"
)

var _ container.Container[string] = (*List[string])(nil)

// IsEmpty will return true when container has no element
func (a *List[T]) IsEmpty() bool {
	return a.size == 0
}

// Size will return the number of elements in the container
func (a *List[T]) Size() int {
	return a.size
}

// Clear the elements inside the container
func (a *List[T]) Clear() {
	a.size = 0
	a.elements = []T{}
}

// Values will return the collection of all element values in the container
func (a *List[T]) Values() []T {
	if a.size == 0 {
		return []T{}
	}

	elements := make([]T, a.size)
	copy(elements, a.elements[:a.size])
	return elements
}

// String returns a string representation of array list
func (a *List[T]) String() string {
	valueStr := make([]string, 0, a.size)
	for _, v := range a.Values() {
		valueStr = append(valueStr, fmt.Sprintf("%v", v))
	}

	return fmt.Sprintf("ArrayList[%s]", strings.Join(valueStr, ", "))
}
