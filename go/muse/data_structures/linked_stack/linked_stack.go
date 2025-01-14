package linked_stack

import "errors"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedStack[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func New[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{head: nil, tail: nil, length: 0}
}

func (stack *LinkedStack[T]) Size() int {
	return stack.length
}

func (stack *LinkedStack[T]) IsEmpty() bool {
	return stack.length == 0
}

func (stack *LinkedStack[T]) Peek() T {
	if stack.length == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	return stack.tail.data
}

func (stack *LinkedStack[T]) Push(element T) {
	node := &Node[T]{data: element, next: nil}

	if stack.length == 0 {
		stack.head = node
	} else {
		stack.tail.next = node
	}

	stack.tail = node

	stack.length++
}

func (stack *LinkedStack[T]) Pop() {
	if stack.length == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	target := stack.tail

	if stack.length == 1 {
		stack.head = nil
		stack.tail = nil
	} else {
		cursor := stack.head
		for range stack.length - 2 {
			cursor = cursor.next
		}
		cursor.next = nil
		stack.tail = cursor
	}

	target.data = *new(T)

	stack.length--
}
