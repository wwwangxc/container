package set

import (
	"github.com/wwwangxc/container"
	"github.com/wwwangxc/container/set/hashset"
)

// NewHashSet return implement of the set by a hash map
func NewHashSet[T comparable](elements ...T) container.Set[T] {
	return hashset.New(elements...)
}
