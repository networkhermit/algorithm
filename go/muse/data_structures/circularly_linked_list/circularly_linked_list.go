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

func (list *CircularlyLinkedList[T]) Size() int {
	return list.length
}

func (list *CircularlyLinkedList[T]) IsEmpty() bool {
	return list.length == 0
}

func (list *CircularlyLinkedList[T]) Get(index int) T {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	cursor := list.tail

	if index != list.length-1 {
		for range index + 1 {
			cursor = cursor.next
		}
	}

	return cursor.data
}

func (list *CircularlyLinkedList[T]) Set(index int, element T) {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	cursor := list.tail

	if index != list.length-1 {
		for range index + 1 {
			cursor = cursor.next
		}
	}

	cursor.data = element
}

func (list *CircularlyLinkedList[T]) Insert(index int, element T) {
	if index < 0 || index > list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	node := &Node[T]{data: element, next: nil}

	switch index {
	case 0:
		if list.length == 0 {
			node.next = node
			list.tail = node
		} else {
			node.next = list.tail.next
			list.tail.next = node
		}
	case list.length:
		node.next = list.tail.next
		list.tail.next = node
		list.tail = node
	default:
		cursor := list.tail
		for range index {
			cursor = cursor.next
		}
		node.next = cursor.next
		cursor.next = node
	}

	list.length++
}

func (list *CircularlyLinkedList[T]) Remove(index int) {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var target *Node[T]

	if index == 0 {
		target = list.tail.next
		if list.length == 1 {
			list.tail = nil
		} else {
			list.tail.next = target.next
		}
	} else {
		cursor := list.tail
		for range index {
			cursor = cursor.next
		}
		target = cursor.next
		cursor.next = target.next
		if index == list.length-1 {
			list.tail = cursor
		}
	}

	target.data = *new(T)

	list.length--
}

func (list *CircularlyLinkedList[T]) Front() T {
	return list.Get(0)
}

func (list *CircularlyLinkedList[T]) Back() T {
	return list.Get(list.length - 1)
}

func (list *CircularlyLinkedList[T]) Prepend(element T) {
	list.Insert(0, element)
}

func (list *CircularlyLinkedList[T]) Append(element T) {
	list.Insert(list.length, element)
}

func (list *CircularlyLinkedList[T]) Poll() {
	list.Remove(0)
}

func (list *CircularlyLinkedList[T]) Eject() {
	list.Remove(list.length - 1)
}
