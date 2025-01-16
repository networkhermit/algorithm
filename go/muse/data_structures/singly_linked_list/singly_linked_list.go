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

func (l *SinglyLinkedList[T]) Size() int {
	return l.length
}

func (l *SinglyLinkedList[T]) IsEmpty() bool {
	return l.length == 0
}

func (l *SinglyLinkedList[T]) Get(index int) T {
	if index < 0 || index >= l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	cursor := l.head

	for range index {
		cursor = cursor.next
	}

	return cursor.data
}

func (l *SinglyLinkedList[T]) Set(index int, element T) {
	if index < 0 || index >= l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	cursor := l.head

	for range index {
		cursor = cursor.next
	}

	cursor.data = element
}

func (l *SinglyLinkedList[T]) Insert(index int, element T) {
	if index < 0 || index > l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	node := &Node[T]{data: element, next: nil}

	if index == 0 {
		if l.length != 0 {
			node.next = l.head
		}
		l.head = node
	} else {
		cursor := l.head
		for range index - 1 {
			cursor = cursor.next
		}
		node.next = cursor.next
		cursor.next = node
	}

	l.length++
}

func (l *SinglyLinkedList[T]) Remove(index int) {
	if index < 0 || index >= l.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var target *Node[T]

	if index == 0 {
		target = l.head
		if l.length == 1 {
			l.head = nil
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
	}

	target.data = *new(T)

	l.length--
}

func (l *SinglyLinkedList[T]) Front() T {
	return l.Get(0)
}

func (l *SinglyLinkedList[T]) Back() T {
	return l.Get(l.length - 1)
}

func (l *SinglyLinkedList[T]) Prepend(element T) {
	l.Insert(0, element)
}

func (l *SinglyLinkedList[T]) Append(element T) {
	l.Insert(l.length, element)
}

func (l *SinglyLinkedList[T]) Poll() {
	l.Remove(0)
}

func (l *SinglyLinkedList[T]) Eject() {
	l.Remove(l.length - 1)
}
