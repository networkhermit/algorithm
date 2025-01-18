package array_queue

import "errors"

const DEFAULT_CAPACITY int = 64

type ArrayQueue[T any] struct {
	data         []T
	front        int
	logicalSize  int
	physicalSize int
}

func New[T any]() *ArrayQueue[T] {
	return NewWithCapacity[T](0)
}

func NewWithCapacity[T any](physicalSize int) *ArrayQueue[T] {
	queue := &ArrayQueue[T]{data: nil, front: 0, logicalSize: 0, physicalSize: DEFAULT_CAPACITY}

	if physicalSize > 1 {
		queue.physicalSize = physicalSize
	}
	queue.data = make([]T, queue.physicalSize)

	return queue
}

func (q *ArrayQueue[T]) Size() int {
	return q.logicalSize
}

func (q *ArrayQueue[T]) IsEmpty() bool {
	return q.logicalSize == 0
}

func (q *ArrayQueue[T]) Peek() T {
	if q.logicalSize == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	return q.data[q.front]
}

func (q *ArrayQueue[T]) Enqueue(element T) {
	if q.logicalSize == q.physicalSize {
		newCapacity := DEFAULT_CAPACITY

		if q.physicalSize >= DEFAULT_CAPACITY {
			newCapacity = q.physicalSize + (q.physicalSize >> 1)
		}

		temp := make([]T, newCapacity)

		cursor := q.front

		for i := range q.logicalSize {
			if cursor == q.physicalSize {
				cursor = 0
			}
			temp[i] = q.data[cursor]
			cursor++
		}

		q.data = temp
		q.front = 0
		q.physicalSize = newCapacity
	}

	q.data[(q.front+q.logicalSize)%q.physicalSize] = element

	q.logicalSize++
}

func (q *ArrayQueue[T]) Dequeue() {
	if q.logicalSize == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	q.data[q.front] = *new(T)

	q.front = (q.front + 1) % q.physicalSize

	q.logicalSize--
}

func (q *ArrayQueue[T]) Capacity() int {
	return q.physicalSize
}

func (q *ArrayQueue[T]) Shrink() {
	temp := make([]T, q.logicalSize)

	cursor := q.front

	for i := range q.logicalSize {
		if cursor == q.physicalSize {
			cursor = 0
		}
		temp[i] = q.data[cursor]
		cursor++
	}

	q.data = temp
	q.front = 0
	q.physicalSize = q.logicalSize
}
