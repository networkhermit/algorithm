package singly_linked_list

import "errors"

type Node[T any] struct {
	data T
	next *Node[T]
}

type SinglyLinkedList[T any] struct {
	head   *Node[T]
	length int
}

func New[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{head: nil, length: 0}
}

func (list *SinglyLinkedList[T]) Size() int {
	return list.length
}

func (list *SinglyLinkedList[T]) IsEmpty() bool {
	return list.length == 0
}

func (list *SinglyLinkedList[T]) Get(index int) T {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	cursor := list.head

	for range index {
		cursor = cursor.next
	}

	return cursor.data
}

func (list *SinglyLinkedList[T]) Set(index int, element T) {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	cursor := list.head

	for range index {
		cursor = cursor.next
	}

	cursor.data = element
}

func (list *SinglyLinkedList[T]) Insert(index int, element T) {
	if index < 0 || index > list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	node := &Node[T]{data: element, next: nil}

	if index == 0 {
		if list.length != 0 {
			node.next = list.head
		}
		list.head = node
	} else {
		cursor := list.head
		for range index - 1 {
			cursor = cursor.next
		}
		node.next = cursor.next
		cursor.next = node
	}

	list.length++
}

func (list *SinglyLinkedList[T]) Remove(index int) {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var target *Node[T]

	if index == 0 {
		target = list.head
		if list.length == 1 {
			list.head = nil
		} else {
			list.head = target.next
		}
	} else {
		cursor := list.head
		for range index - 1 {
			cursor = cursor.next
		}
		target = cursor.next
		cursor.next = target.next
	}

	target.data = *new(T)

	list.length--
}

func (list *SinglyLinkedList[T]) Front() T {
	return list.Get(0)
}

func (list *SinglyLinkedList[T]) Back() T {
	return list.Get(list.length - 1)
}

func (list *SinglyLinkedList[T]) Prepend(element T) {
	list.Insert(0, element)
}

func (list *SinglyLinkedList[T]) Append(element T) {
	list.Insert(list.length, element)
}

func (list *SinglyLinkedList[T]) Poll() {
	list.Remove(0)
}

func (list *SinglyLinkedList[T]) Eject() {
	list.Remove(list.length - 1)
}
