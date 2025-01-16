package linked_stack

import "errors"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedStack[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func New[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{head: nil, tail: nil, length: 0}
}

func (s *LinkedStack[T]) Size() int {
	return s.length
}

func (s *LinkedStack[T]) IsEmpty() bool {
	return s.length == 0
}

func (s *LinkedStack[T]) Peek() T {
	if s.length == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	return s.tail.data
}

func (s *LinkedStack[T]) Push(element T) {
	node := &Node[T]{data: element, next: nil}

	if s.length == 0 {
		s.head = node
	} else {
		s.tail.next = node
	}

	s.tail = node

	s.length++
}

func (s *LinkedStack[T]) Pop() {
	if s.length == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	target := s.tail

	if s.length == 1 {
		s.head = nil
		s.tail = nil
	} else {
		cursor := s.head
		for range s.length - 2 {
			cursor = cursor.next
		}
		cursor.next = nil
		s.tail = cursor
	}

	target.data = *new(T)

	s.length--
}
