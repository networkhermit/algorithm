package circularly_linked_list

import "errors"

type Node[T any] struct {
	data T
	next *Node[T]
}

type CircularlyLinkedList[T any] struct {
	tail   *Node[T]
	length int
}

func New[T any]() *CircularlyLinkedList[T] {
	return &CircularlyLinkedList[T]{tail: nil, length: 0}
}

func (l *CircularlyLinkedList[T]) Size() int {
	return l.length
}

func (l *CircularlyLinkedList[T]) IsEmpty() bool {
	return l.length == 0
}

func (l *CircularlyLinkedList[T]) Get(index int) T {
	if index < 0 || index >= l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	cursor := l.tail

	if index != l.length-1 {
		for range index + 1 {
			cursor = cursor.next
		}
	}

	return cursor.data
}

func (l *CircularlyLinkedList[T]) Set(index int, element T) {
	if index < 0 || index >= l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	cursor := l.tail

	if index != l.length-1 {
		for range index + 1 {
			cursor = cursor.next
		}
	}

	cursor.data = element
}

func (l *CircularlyLinkedList[T]) Insert(index int, element T) {
	if index < 0 || index > l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	node := &Node[T]{data: element, next: nil}

	switch index {
	case 0:
		if l.length == 0 {
			node.next = node
			l.tail = node
		} else {
			node.next = l.tail.next
			l.tail.next = node
		}
	case l.length:
		node.next = l.tail.next
		l.tail.next = node
		l.tail = node
	default:
		cursor := l.tail
		for range index {
			cursor = cursor.next
		}
		node.next = cursor.next
		cursor.next = node
	}

	l.length++
}

func (l *CircularlyLinkedList[T]) Remove(index int) {
	if index < 0 || index >= l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var target *Node[T]

	if index == 0 {
		target = l.tail.next
		if l.length == 1 {
			l.tail = nil
		} else {
			l.tail.next = target.next
		}
	} else {
		cursor := l.tail
		for range index {
			cursor = cursor.next
		}
		target = cursor.next
		cursor.next = target.next
		if index == l.length-1 {
			l.tail = cursor
		}
	}

	target.data = *new(T)

	l.length--
}

func (l *CircularlyLinkedList[T]) Front() T {
	return l.Get(0)
}

func (l *CircularlyLinkedList[T]) Back() T {
	return l.Get(l.length - 1)
}

func (l *CircularlyLinkedList[T]) Prepend(element T) {
	l.Insert(0, element)
}

func (l *CircularlyLinkedList[T]) Append(element T) {
	l.Insert(l.length, element)
}

func (l *CircularlyLinkedList[T]) Poll() {
	l.Remove(0)
}

func (l *CircularlyLinkedList[T]) Eject() {
	l.Remove(l.length - 1)
}
