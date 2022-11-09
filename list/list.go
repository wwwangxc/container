package list

import (
	"github.com/wwwangxc/container"
	"github.com/wwwangxc/container/list/arraylist"
	"github.com/wwwangxc/container/list/doublylinkedlist"
	"github.com/wwwangxc/container/list/singlylinkedlist"
)

// NewArray return implement of the array list
func NewArray[T comparable](values ...T) container.List[T] {
	return arraylist.New(values...)
}

// NewSingly return implement of the singly linked list
func NewSingly[T comparable](values ...T) container.List[T] {
	return singlylinkedlist.New(values...)
}

// NewDoubly return implement of the doubly linked list
func NewDoubly[T comparable](values ...T) container.List[T] {
	return doublylinkedlist.New(values...)
}
