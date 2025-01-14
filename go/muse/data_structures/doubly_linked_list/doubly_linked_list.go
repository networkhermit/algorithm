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

func (list *DoublyLinkedList[T]) Size() int {
	return list.length
}

func (list *DoublyLinkedList[T]) IsEmpty() bool {
	return list.length == 0
}

func (list *DoublyLinkedList[T]) Get(index int) T {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var cursor *Node[T]

	if index < list.length>>1 {
		cursor = list.head
		for range index {
			cursor = cursor.next
		}
	} else {
		cursor = list.tail
		for i := list.length - 1; i > index; i-- {
			cursor = cursor.prev
		}
	}

	return cursor.data
}

func (list *DoublyLinkedList[T]) Set(index int, element T) {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var cursor *Node[T]

	if index < list.length>>1 {
		cursor = list.head
		for range index {
			cursor = cursor.next
		}
	} else {
		cursor = list.tail
		for i := list.length - 1; i > index; i-- {
			cursor = cursor.prev
		}
	}

	cursor.data = element
}

func (list *DoublyLinkedList[T]) Insert(index int, element T) {
	if index < 0 || index > list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	node := &Node[T]{data: element, next: nil, prev: nil}

	switch index {
	case 0:
		if list.length != 0 {
			node.next = list.head
			list.head.prev = node
		} else {
			list.tail = node
		}
		list.head = node
	case list.length:
		node.prev = list.tail
		list.tail.next = node
		list.tail = node
	default:
		var cursor *Node[T]
		if index < list.length>>1 {
			cursor = list.head
			for range index {
				cursor = cursor.next
			}
		} else {
			cursor = list.tail
			for i := list.length - 1; i > index; i-- {
				cursor = cursor.prev
			}
		}
		node.next = cursor
		node.prev = cursor.prev
		cursor.prev.next = node
		cursor.prev = node
	}

	list.length++
}

func (list *DoublyLinkedList[T]) Remove(index int) {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var target *Node[T]

	switch index {
	case 0:
		target = list.head
		if list.length == 1 {
			list.head = nil
			list.tail = nil
		} else {
			target.next.prev = nil
			list.head = target.next
		}
	case list.length - 1:
		target = list.tail
		target.prev.next = nil
		list.tail = target.prev
	default:
		if index < list.length>>1 {
			target = list.head
			for range index {
				target = target.next
			}
		} else {
			target = list.tail
			for i := list.length - 1; i > index; i-- {
				target = target.prev
			}
		}
		target.prev.next = target.next
		target.next.prev = target.prev
	}

	target.data = *new(T)

	list.length--
}

func (list *DoublyLinkedList[T]) Front() T {
	return list.Get(0)
}

func (list *DoublyLinkedList[T]) Back() T {
	return list.Get(list.length - 1)
}

func (list *DoublyLinkedList[T]) Prepend(element T) {
	list.Insert(0, element)
}

func (list *DoublyLinkedList[T]) Append(element T) {
	list.Insert(list.length, element)
}

func (list *DoublyLinkedList[T]) Poll() {
	list.Remove(0)
}

func (list *DoublyLinkedList[T]) Eject() {
	list.Remove(list.length - 1)
}
