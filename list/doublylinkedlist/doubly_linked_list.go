package doublylinkedlist

// List implement of the doubly linked list
// Not thread safety
// Reference: https://en.wikipedia.org/wiki/Doubly_linked_list
type List[T comparable] struct {
	head *element[T]
	tail *element[T]
	size int
}

type element[T comparable] struct {
	value T
	next  *element[T]
	prev  *element[T]
}

// New implement of the doubly linked list
func New[T comparable](values ...T) *List[T] {
	l := &List[T]{}
	l.Append(values...)
	return l
}

// outOfRange check if index is out of range
func (l *List[T]) outOfRange(index int) bool {
	return index < 0 || index >= l.size
}

// append value into element list
func (l *List[T]) append(value T) {
	e := &element[T]{value: value}

	if l.size == 0 {
		l.head = e
		l.tail = e
	} else {
		l.tail.next = e
		e.prev = l.tail
		l.tail = e
	}

	l.size++
}
