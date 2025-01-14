package array_stack

import "errors"

const DEFAULT_CAPACITY int = 64

type ArrayStack[T any] struct {
	data         []T
	logicalSize  int
	physicalSize int
}

func New[T any]() *ArrayStack[T] {
	return NewWithCapacity[T](0)
}

func NewWithCapacity[T any](physicalSize int) *ArrayStack[T] {
	stack := &ArrayStack[T]{data: nil, logicalSize: 0, physicalSize: DEFAULT_CAPACITY}

	if physicalSize > 1 {
		stack.physicalSize = physicalSize
	}
	stack.data = make([]T, stack.physicalSize)

	return stack
}

func (stack *ArrayStack[T]) Size() int {
	return stack.logicalSize
}

func (stack *ArrayStack[T]) IsEmpty() bool {
	return stack.logicalSize == 0
}

func (stack *ArrayStack[T]) Peek() T {
	if stack.logicalSize == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	return stack.data[stack.logicalSize-1]
}

func (stack *ArrayStack[T]) Push(element T) {
	if stack.logicalSize == stack.physicalSize {
		newCapacity := DEFAULT_CAPACITY

		if stack.physicalSize >= DEFAULT_CAPACITY {
			newCapacity = stack.physicalSize + (stack.physicalSize >> 1)
		}

		temp := make([]T, newCapacity)

		for i := range stack.logicalSize {
			temp[i] = stack.data[i]
		}

		stack.data = temp
		stack.physicalSize = newCapacity
	}

	stack.data[stack.logicalSize] = element

	stack.logicalSize++
}

func (stack *ArrayStack[T]) Pop() {
	if stack.logicalSize == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	stack.logicalSize--

	stack.data[stack.logicalSize] = *new(T)
}

func (stack *ArrayStack[T]) Capacity() int {
	return stack.physicalSize
}

func (stack *ArrayStack[T]) Shrink() {
	temp := make([]T, stack.logicalSize)

	for i := range stack.logicalSize {
		temp[i] = stack.data[i]
	}

	stack.data = temp
	stack.physicalSize = stack.logicalSize
}
