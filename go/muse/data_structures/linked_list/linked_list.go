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

func (list *LinkedList[T]) Size() int {
	return list.length
}

func (list *LinkedList[T]) IsEmpty() bool {
	return list.length == 0
}

func (list *LinkedList[T]) Get(index int) T {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var cursor *Node[T]

	if index == list.length-1 {
		cursor = list.tail
	} else {
		cursor = list.head
		for range index {
			cursor = cursor.next
		}
	}

	return cursor.data
}

func (list *LinkedList[T]) Set(index int, element T) {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var cursor *Node[T]

	if index == list.length-1 {
		cursor = list.tail
	} else {
		cursor = list.head
		for range index {
			cursor = cursor.next
		}
	}

	cursor.data = element
}

func (list *LinkedList[T]) Insert(index int, element T) {
	if index < 0 || index > list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	node := &Node[T]{data: element, next: nil}

	switch index {
	case 0:
		if list.length != 0 {
			node.next = list.head
		} else {
			list.tail = node
		}
		list.head = node
	case list.length:
		list.tail.next = node
		list.tail = node
	default:
		cursor := list.head
		for i, bound := 0, index-1; i < bound; i++ {
			cursor = cursor.next
		}
		node.next = cursor.next
		cursor.next = node
	}

	list.length++
}

func (list *LinkedList[T]) Remove(index int) {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var target *Node[T]

	if index == 0 {
		target = list.head
		if list.length == 1 {
			list.head = nil
			list.tail = nil
		} else {
			list.head = target.next
		}
	} else {
		cursor := list.head
		for i, bound := 0, index-1; i < bound; i++ {
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

func (list *LinkedList[T]) Front() T {
	return list.Get(0)
}

func (list *LinkedList[T]) Back() T {
	return list.Get(list.length - 1)
}

func (list *LinkedList[T]) Prepend(element T) {
	list.Insert(0, element)
}

func (list *LinkedList[T]) Append(element T) {
	list.Insert(list.length, element)
}

func (list *LinkedList[T]) Poll() {
	list.Remove(0)
}

func (list *LinkedList[T]) Eject() {
	list.Remove(list.length - 1)
}
