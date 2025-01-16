package doubly_linked_list

import "errors"

type Node[T any] struct {
	data T
	next *Node[T]
	prev *Node[T]
}

type DoublyLinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func New[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{head: nil, tail: nil, length: 0}
}

func (l *DoublyLinkedList[T]) Size() int {
	return l.length
}

func (l *DoublyLinkedList[T]) IsEmpty() bool {
	return l.length == 0
}

func (l *DoublyLinkedList[T]) Get(index int) T {
	if index < 0 || index >= l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var cursor *Node[T]

	if index < l.length>>1 {
		cursor = l.head
		for range index {
			cursor = cursor.next
		}
	} else {
		cursor = l.tail
		for i := l.length - 1; i > index; i-- {
			cursor = cursor.prev
		}
	}

	return cursor.data
}

func (l *DoublyLinkedList[T]) Set(index int, element T) {
	if index < 0 || index >= l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var cursor *Node[T]

	if index < l.length>>1 {
		cursor = l.head
		for range index {
			cursor = cursor.next
		}
	} else {
		cursor = l.tail
		for i := l.length - 1; i > index; i-- {
			cursor = cursor.prev
		}
	}

	cursor.data = element
}

func (l *DoublyLinkedList[T]) Insert(index int, element T) {
	if index < 0 || index > l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	node := &Node[T]{data: element, next: nil, prev: nil}

	switch index {
	case 0:
		if l.length != 0 {
			node.next = l.head
			l.head.prev = node
		} else {
			l.tail = node
		}
		l.head = node
	case l.length:
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	default:
		var cursor *Node[T]
		if index < l.length>>1 {
			cursor = l.head
			for range index {
				cursor = cursor.next
			}
		} else {
			cursor = l.tail
			for i := l.length - 1; i > index; i-- {
				cursor = cursor.prev
			}
		}
		node.next = cursor
		node.prev = cursor.prev
		cursor.prev.next = node
		cursor.prev = node
	}

	l.length++
}

func (l *DoublyLinkedList[T]) Remove(index int) {
	if index < 0 || index >= l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var target *Node[T]

	switch index {
	case 0:
		target = l.head
		if l.length == 1 {
			l.head = nil
			l.tail = nil
		} else {
			target.next.prev = nil
			l.head = target.next
		}
	case l.length - 1:
		target = l.tail
		target.prev.next = nil
		l.tail = target.prev
	default:
		if index < l.length>>1 {
			target = l.head
			for range index {
				target = target.next
			}
		} else {
			target = l.tail
			for i := l.length - 1; i > index; i-- {
				target = target.prev
			}
		}
		target.prev.next = target.next
		target.next.prev = target.prev
	}

	target.data = *new(T)

	l.length--
}

func (l *DoublyLinkedList[T]) Front() T {
	return l.Get(0)
}

func (l *DoublyLinkedList[T]) Back() T {
	return l.Get(l.length - 1)
}

func (l *DoublyLinkedList[T]) Prepend(element T) {
	l.Insert(0, element)
}

func (l *DoublyLinkedList[T]) Append(element T) {
	l.Insert(l.length, element)
}

func (l *DoublyLinkedList[T]) Poll() {
	l.Remove(0)
}

func (l *DoublyLinkedList[T]) Eject() {
	l.Remove(l.length - 1)
}
