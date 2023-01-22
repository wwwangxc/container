package maps

import "github.com/wwwangxc/container/maps/linkedmap"

// NewLinkedMap return linked map implement sorted associative array
func NewLinkedMap[K comparable, V any]() *linkedmap.Map[K, V] {
	return linkedmap.New[K, V]()
}
