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
func (l *List[T]) append(value T) {
	index := l.size
	if index < len(l.elements) {
		l.elements[index] = value
	} else {
		l.elements = append(l.elements, value)
	}
	l.size++
}

// outOfRange check if index is out of range
func (l *List[T]) outOfRange(index int) bool {
	return index < 0 || index >= l.size
}

// tryShrink shrink the array if necessary, i.e. when size less than shrink threshold of current capacity
func (l *List[T]) tryShrink() {
	referSize := int(float32(cap(l.elements)) * thresholdShrink)
	if l.size < referSize {
		l.resize(referSize)
	}
}

// resize current list
func (l *List[T]) resize(capacity int) {
	elements := make([]T, l.size, capacity)
	copy(elements, l.elements[:l.size])
	l.elements = elements
}
