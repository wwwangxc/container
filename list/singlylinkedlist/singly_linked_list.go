package singlylinkedlist

// List implement of the singly linked list
// Not thread safety
// Reference: https://en.wikipedia.org/wiki/Linked_list#Singly_linked_list
type List[T comparable] struct {
	head *element[T]
	tail *element[T]
	size int
}

type element[T comparable] struct {
	value T
	next  *element[T]
}

// New implement of the singly linked list
func New[T comparable](values ...T) *List[T] {
	l := &List[T]{}
	l.Append(values...)
	return l
}

// outOfRange check if index is out of range
func (a *List[T]) outOfRange(index int) bool {
	return index < 0 || index >= a.size
}

// append value into element list
func (a *List[T]) append(value T) {
	e := &element[T]{value: value}

	if a.size == 0 {
		a.head = e
		a.tail = e
	} else {
		a.tail.next = e
		a.tail = e
	}

	a.size++
}
