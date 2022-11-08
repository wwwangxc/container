package container

// Set is the base feature provided by all set structures
// Reference: https://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
type Set[T comparable] interface {

	// Add all given elements into the set
	Add(elements ...T)

	// Remove all given elements from the set
	Remove(elements ...T)

	// Contains all given elements
	Contains(elements ...T) bool

	// Union returns the union of two sets.
	// All elements of the new set exist in the current set or exist in the target set.
	// Reference: https://en.wikipedia.org/wiki/Union_(set_theory)
	Union(target Set[T]) Set[T]

	// Intersection returns the intersection between two sets.
	// All elements of the new set exist in both current set and target set.
	// Reference: https://en.wikipedia.org/wiki/Intersection_(set_theory)
	Intersection(target Set[T]) Set[T]

	// Difference returns the difference between two sets.
	// All elements of the new set exist in the current set but not in the target set.
	// Reference: https://proofwiki.org/wiki/Definition:Set_Difference
	Difference(target Set[T]) Set[T]

	// Container is the base feature provided by all data structures
	Container[T]
}
