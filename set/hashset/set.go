package hashset

import "github.com/wwwangxc/container"

var _ container.Set[int] = (*Set[int])(nil)

// Add all given elements into the set
func (s *Set[T]) Add(elements ...T) {
	if s.m == nil {
		s.m = make(map[T]struct{})
	}

	for _, v := range elements {
		s.m[v] = empty
	}
}

// Remove all given elements from the set
func (s *Set[T]) Remove(elements ...T) {
	if s.IsEmpty() {
		return
	}

	for _, v := range elements {
		delete(s.m, v)
	}
}

// Contains all given elements
func (s *Set[T]) Contains(elements ...T) bool {
	if s.IsEmpty() {
		return false
	}

	for _, v := range elements {
		if _, ok := s.m[v]; !ok {
			return false
		}
	}

	return true
}

// Union returns the union of two sets.
// All elements of the new set exist in the current set or exist in the target set.
// Reference: https://en.wikipedia.org/wiki/Union_(set_theory)
func (s *Set[T]) Union(target container.Set[T]) container.Set[T] {
	if target == nil {
		return New(s.Values()...)
	}

	set := New(s.Values()...)
	for _, v := range target.Values() {
		set.Add(v)
	}

	return set
}

// Intersection returns the intersection between two sets.
// All elements of the new set exist in both current set and target set.
// Reference: https://en.wikipedia.org/wiki/Intersection_(set_theory)
func (s *Set[T]) Intersection(target container.Set[T]) container.Set[T] {
	if target == nil {
		return New[T]()
	}

	set := New[T]()
	minimal := container.Set[T](s)
	if s.Size() > target.Size() {
		minimal = target
	}

	for _, v := range minimal.Values() {
		if s.Contains(v) && target.Contains(v) {
			set.Add(v)
		}
	}

	return set
}

// Difference returns the difference between two sets.
// All elements of the new set exist in the current set but not in the target set.
// Reference: https://proofwiki.org/wiki/Definition:Set_Difference
func (s *Set[T]) Difference(target container.Set[T]) container.Set[T] {
	if target == nil {
		return New(s.Values()...)
	}

	set := New[T]()
	for _, v := range s.Values() {
		if !target.Contains(v) {
			set.Add(v)
		}
	}

	return set
}
