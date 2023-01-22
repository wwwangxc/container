package linkedmap

import (
	"github.com/wwwangxc/container"
	"github.com/wwwangxc/container/list/arraylist"
)

// Map implement of the sorted associative array
// Not thread safety
// Reference: https://en.wikipedia.org/wiki/Associative_array
type Map[K comparable, V any] struct {
	m    map[K]V
	list container.List[K]
}

// New implement sorted associative array
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m:    map[K]V{},
		list: arraylist.New[K](),
	}
}
