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
	queue := &ArrayQueue[T]{data: nil, logicalSize: 0, physicalSize: DEFAULT_CAPACITY}

	if physicalSize > 1 {
		queue.physicalSize = physicalSize
	}
	queue.data = make([]T, queue.physicalSize)

	return queue
}

func (queue *ArrayQueue[T]) Size() int {
	return queue.logicalSize
}

func (queue *ArrayQueue[T]) IsEmpty() bool {
	return queue.logicalSize == 0
}

func (queue *ArrayQueue[T]) Peek() T {
	if queue.logicalSize == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	return queue.data[queue.front]
}

func (queue *ArrayQueue[T]) Enqueue(element T) {
	if queue.logicalSize == queue.physicalSize {
		newCapacity := DEFAULT_CAPACITY

		if queue.physicalSize >= DEFAULT_CAPACITY {
			newCapacity = queue.physicalSize + (queue.physicalSize >> 1)
		}

		temp := make([]T, newCapacity)

		cursor := queue.front

		for i := range queue.logicalSize {
			if cursor == queue.physicalSize {
				cursor = 0
			}
			temp[i] = queue.data[cursor]
			cursor++
		}

		queue.data = temp
		queue.front = 0
		queue.physicalSize = newCapacity
	}

	queue.data[(queue.front+queue.logicalSize)%queue.physicalSize] = element

	queue.logicalSize++
}

func (queue *ArrayQueue[T]) Dequeue() {
	if queue.logicalSize == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	queue.data[queue.front] = *new(T)

	queue.front = (queue.front + 1) % queue.physicalSize

	queue.logicalSize--
}

func (queue *ArrayQueue[T]) Capacity() int {
	return queue.physicalSize
}

func (queue *ArrayQueue[T]) Shrink() {
	temp := make([]T, queue.logicalSize)

	cursor := queue.front

	for i := range queue.logicalSize {
		if cursor == queue.physicalSize {
			cursor = 0
		}
		temp[i] = queue.data[cursor]
		cursor++
	}

	queue.data = temp
	queue.front = 0
	queue.physicalSize = queue.logicalSize
}
