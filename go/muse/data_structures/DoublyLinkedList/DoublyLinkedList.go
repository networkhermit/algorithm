package DoublyLinkedList

import "errors"

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func New() *DoublyLinkedList {
	return &DoublyLinkedList{head: nil, tail: nil, length: 0}
}

func (list *DoublyLinkedList) Size() int {
	return list.length
}

func (list *DoublyLinkedList) IsEmpty() bool {
	return list.length == 0
}

func (list *DoublyLinkedList) Get(index int) int {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var cursor *Node

	if index < list.length>>1 {
		cursor = list.head
		for i := 0; i < index; i++ {
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

func (list *DoublyLinkedList) Set(index int, element int) {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var cursor *Node

	if index < list.length>>1 {
		cursor = list.head
		for i := 0; i < index; i++ {
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

func (list *DoublyLinkedList) Insert(index int, element int) {
	if index < 0 || index > list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	node := &Node{data: element, next: nil, prev: nil}

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
		var cursor *Node
		if index < list.length>>1 {
			cursor = list.head
			for i := 0; i < index; i++ {
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

func (list *DoublyLinkedList) Remove(index int) {
	if index < 0 || index >= list.length {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	var target *Node

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
			for i := 0; i < index; i++ {
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

	target.data = int(0)

	list.length--
}

func (list *DoublyLinkedList) Front() int {
	return list.Get(0)
}

func (list *DoublyLinkedList) Back() int {
	return list.Get(list.length - 1)
}

func (list *DoublyLinkedList) Prepend(element int) {
	list.Insert(0, element)
}

func (list *DoublyLinkedList) Append(element int) {
	list.Insert(list.length, element)
}

func (list *DoublyLinkedList) Poll() {
	list.Remove(0)
}

func (list *DoublyLinkedList) Eject() {
	list.Remove(list.length - 1)
}
