package hashset

import (
	"fmt"
	"strings"

	"github.com/wwwangxc/container"
)

var _ container.Container[int] = (*Set[int])(nil)

// IsEmpty will return true when set has no element
func (s *Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Size will return the number of elements in the set
func (s *Set[T]) Size() int {
	return len(s.m)
}

// Clear the elements inside the set
func (s *Set[T]) Clear() {
	s.m = make(map[T]struct{})
}

// Values will return the collection of all element in the set
func (s *Set[T]) Values() []T {
	if s.IsEmpty() {
		return []T{}
	}

	elements := make([]T, 0, s.Size())
	for v := range s.m {
		elements = append(elements, v)
	}

	return elements
}

// String returns a string representation of set
func (s *Set[T]) String() string {
	valueStr := make([]string, 0, s.Size())
	for _, v := range s.Values() {
		valueStr = append(valueStr, fmt.Sprintf("%v", v))
	}

	return fmt.Sprintf("HashSet[%s]", strings.Join(valueStr, ", "))
}
