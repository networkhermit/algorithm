package linked_list

import "errors"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{head: nil, tail: nil, length: 0}
}

func (l *LinkedList[T]) Size() int {
	return l.length
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.length == 0
}

func (l *LinkedList[T]) Get(index int) T {
	if index < 0 || index >= l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var cursor *Node[T]

	if index == l.length-1 {
		cursor = l.tail
	} else {
		cursor = l.head
		for range index {
			cursor = cursor.next
		}
	}

	return cursor.data
}

func (l *LinkedList[T]) Set(index int, element T) {
	if index < 0 || index >= l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var cursor *Node[T]

	if index == l.length-1 {
		cursor = l.tail
	} else {
		cursor = l.head
		for range index {
			cursor = cursor.next
		}
	}

	cursor.data = element
}

func (l *LinkedList[T]) Insert(index int, element T) {
	if index < 0 || index > l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	node := &Node[T]{data: element, next: nil}

	switch index {
	case 0:
		if l.length != 0 {
			node.next = l.head
		} else {
			l.tail = node
		}
		l.head = node
	case l.length:
		l.tail.next = node
		l.tail = node
	default:
		cursor := l.head
		for range index - 1 {
			cursor = cursor.next
		}
		node.next = cursor.next
		cursor.next = node
	}

	l.length++
}

func (l *LinkedList[T]) Remove(index int) {
	if index < 0 || index >= l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var target *Node[T]

	if index == 0 {
		target = l.head
		if l.length == 1 {
			l.head = nil
			l.tail = nil
		} else {
			l.head = target.next
		}
	} else {
		cursor := l.head
		for range index - 1 {
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

func (l *LinkedList[T]) Front() T {
	return l.Get(0)
}

func (l *LinkedList[T]) Back() T {
	return l.Get(l.length - 1)
}

func (l *LinkedList[T]) Prepend(element T) {
	l.Insert(0, element)
}

func (l *LinkedList[T]) Append(element T) {
	l.Insert(l.length, element)
}

func (l *LinkedList[T]) Poll() {
	l.Remove(0)
}

func (l *LinkedList[T]) Eject() {
	l.Remove(l.length - 1)
}
