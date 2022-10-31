package arraylist

const thresholdShrink float32 = 0.25

// List implement of the array list
// Not thread safety
// Reference: https://en.wikipedia.org/wiki/Array_(data_structure)
type List[T comparable] struct {
	elements []T
	size     int
}

// New implement of the array list
func New[T comparable](values ...T) *List[T] {
	arrayList := &List[T]{
		elements: make([]T, 0, len(values)),
		size:     0,
	}
	arrayList.Append(values...)
	return arrayList
}

// append value into element list
func (a *List[T]) append(value T) {
	index := a.size
	if index < len(a.elements) {
		a.elements[index] = value
	} else {
		a.elements = append(a.elements, value)
	}
	a.size++
}

// outOfRange check if index is out of range
func (a *List[T]) outOfRange(index int) bool {
	return index < 0 || index >= a.size
}

// tryShrink shrink the array if necessary, i.e. when size less than shrink threshold of current capacity
func (a *List[T]) tryShrink() {
	referSize := int(float32(cap(a.elements)) * thresholdShrink)
	if a.size < referSize {
		a.resize(referSize)
	}
}

// resize current list
func (a *List[T]) resize(capacity int) {
	elements := make([]T, a.size, capacity)
	copy(elements, a.elements[:a.size])
	a.elements = elements
}
