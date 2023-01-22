package linkedmap

import (
	"fmt"
	"strings"

	"github.com/wwwangxc/container"
	"github.com/wwwangxc/container/list"
)

var _ container.Container[string] = (*Map[string, string])(nil)

// IsEmpty will return true when container has no element
func (m *Map[K, V]) IsEmpty() bool {
	return m == nil || len(m.m) == 0
}

// Size will return the number of elements in the container
func (m *Map[K, V]) Size() int {
	if m == nil {
		return 0
	}

	return len(m.m)
}

// Clear the elements inside the container
func (m *Map[K, V]) Clear() {
	m.m = map[K]V{}
	m.list = list.NewArray[K]()
}

// Values will return the collection of all element values in the container
func (m *Map[K, V]) Values() []V {
	if m.IsEmpty() {
		return []V{}
	}

	values := make([]V, 0, len(m.m))
	m.list.Each(func(_ int, k K) bool {
		values = append(values, m.m[k])
		return true
	})

	return values
}

// String returns a string representation of linked map
func (m *Map[K, V]) String() string {
	valueStr := make([]string, 0, len(m.m))
	m.list.Each(func(_ int, k K) bool {
		valueStr = append(valueStr, fmt.Sprintf("%v=>%v", k, m.m[k]))
		return true
	})

	return fmt.Sprintf("LinkedMap[%s]", strings.Join(valueStr, ", "))
}
