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

func (s *ArrayStack[T]) Size() int {
	return s.logicalSize
}

func (s *ArrayStack[T]) IsEmpty() bool {
	return s.logicalSize == 0
}

func (s *ArrayStack[T]) Peek() T {
	if s.logicalSize == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	return s.data[s.logicalSize-1]
}

func (s *ArrayStack[T]) Push(element T) {
	if s.logicalSize == s.physicalSize {
		newCapacity := DEFAULT_CAPACITY

		if s.physicalSize >= DEFAULT_CAPACITY {
			newCapacity = s.physicalSize + (s.physicalSize >> 1)
		}

		temp := make([]T, newCapacity)

		for i := range s.logicalSize {
			temp[i] = s.data[i]
		}

		s.data = temp
		s.physicalSize = newCapacity
	}

	s.data[s.logicalSize] = element

	s.logicalSize++
}

func (s *ArrayStack[T]) Pop() {
	if s.logicalSize == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	s.logicalSize--

	s.data[s.logicalSize] = *new(T)
}

func (s *ArrayStack[T]) Capacity() int {
	return s.physicalSize
}

func (s *ArrayStack[T]) Shrink() {
	temp := make([]T, s.logicalSize)

	for i := range s.logicalSize {
		temp[i] = s.data[i]
	}

	s.data = temp
	s.physicalSize = s.logicalSize
}
