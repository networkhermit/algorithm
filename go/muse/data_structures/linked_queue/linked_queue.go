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

func (q *LinkedQueue[T]) Size() int {
	return q.length
}

func (q *LinkedQueue[T]) IsEmpty() bool {
	return q.length == 0
}

func (q *LinkedQueue[T]) Peek() T {
	if q.length == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	return q.head.data
}

func (q *LinkedQueue[T]) Enqueue(element T) {
	node := &Node[T]{data: element, next: nil}

	if q.length == 0 {
		q.head = node
	} else {
		q.tail.next = node
	}

	q.tail = node

	q.length++
}

func (q *LinkedQueue[T]) Dequeue() {
	if q.length == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	target := q.head

	if q.length == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = target.next
	}

	target.data = *new(T)

	q.length--
}
