package circularly_linked_list

import "errors"

type Node struct {
	data int
	next *Node
}

type CircularlyLinkedList struct {
	tail   *Node
	length int
}

func New() *CircularlyLinkedList {
	return &CircularlyLinkedList{tail: nil, length: 0}
}

func (list *CircularlyLinkedList) Size() int {
	return list.length
}

func (list *CircularlyLinkedList) IsEmpty() bool {
	return list.length == 0
}

func (list *CircularlyLinkedList) Get(index int) int {
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

func (list *CircularlyLinkedList) Set(index int, element int) {
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

func (list *CircularlyLinkedList) Insert(index int, element int) {
	if index < 0 || index > list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	node := &Node{data: element, next: nil}

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
		for i, bound := 0, index-1; i <= bound; i++ {
			cursor = cursor.next
		}
		node.next = cursor.next
		cursor.next = node
	}

	list.length++
}

func (list *CircularlyLinkedList) Remove(index int) {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var target *Node

	if index == 0 {
		target = list.tail.next
		if list.length == 1 {
			list.tail = nil
		} else {
			list.tail.next = target.next
		}
	} else {
		cursor := list.tail
		for i, bound := 0, index-1; i <= bound; i++ {
			cursor = cursor.next
		}
		target = cursor.next
		cursor.next = target.next
		if index == list.length-1 {
			list.tail = cursor
		}
	}

	target.data = int(0)

	list.length--
}

func (list *CircularlyLinkedList) Front() int {
	return list.Get(0)
}

func (list *CircularlyLinkedList) Back() int {
	return list.Get(list.length - 1)
}

func (list *CircularlyLinkedList) Prepend(element int) {
	list.Insert(0, element)
}

func (list *CircularlyLinkedList) Append(element int) {
	list.Insert(list.length, element)
}

func (list *CircularlyLinkedList) Poll() {
	list.Remove(0)
}

func (list *CircularlyLinkedList) Eject() {
	list.Remove(list.length - 1)
}
