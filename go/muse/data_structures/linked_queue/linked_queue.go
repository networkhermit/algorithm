package linked_queue

import "errors"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedQueue[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func New[T any]() *LinkedQueue[T] {
	return &LinkedQueue[T]{head: nil, tail: nil, length: 0}
}

func (queue *LinkedQueue[T]) Size() int {
	return queue.length
}

func (queue *LinkedQueue[T]) IsEmpty() bool {
	return queue.length == 0
}

func (queue *LinkedQueue[T]) Peek() T {
	if queue.length == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	return queue.head.data
}

func (queue *LinkedQueue[T]) Enqueue(element T) {
	node := &Node[T]{data: element, next: nil}

	if queue.length == 0 {
		queue.head = node
	} else {
		queue.tail.next = node
	}

	queue.tail = node

	queue.length++
}

func (queue *LinkedQueue[T]) Dequeue() {
	if queue.length == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	target := queue.head

	if queue.length == 1 {
		queue.head = nil
		queue.tail = nil
	} else {
		queue.head = target.next
	}

	target.data = *new(T)

	queue.length--
}
