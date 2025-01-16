package array_list

import "errors"

const DEFAULT_CAPACITY int = 64

type ArrayList[T any] struct {
	data         []T
	logicalSize  int
	physicalSize int
}

func New[T any]() *ArrayList[T] {
	return NewWithCapacity[T](0)
}

func NewWithCapacity[T any](physicalSize int) *ArrayList[T] {
	list := &ArrayList[T]{data: nil, logicalSize: 0, physicalSize: DEFAULT_CAPACITY}

	if physicalSize > 1 {
		list.physicalSize = physicalSize
	}
	list.data = make([]T, list.physicalSize)

	return list
}

func (l *ArrayList[T]) Size() int {
	return l.logicalSize
}

func (l *ArrayList[T]) IsEmpty() bool {
	return l.logicalSize == 0
}

func (l *ArrayList[T]) Get(index int) T {
	if index < 0 || index >= l.logicalSize {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	return l.data[index]
}

func (l *ArrayList[T]) Set(index int, element T) {
	if index < 0 || index >= l.logicalSize {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	l.data[index] = element
}

func (l *ArrayList[T]) Insert(index int, element T) {
	if index < 0 || index > l.logicalSize {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	if l.logicalSize == l.physicalSize {
		newCapacity := DEFAULT_CAPACITY

		if l.physicalSize >= DEFAULT_CAPACITY {
			newCapacity = l.physicalSize + (l.physicalSize >> 1)
		}

		temp := make([]T, newCapacity)

		for i := range l.logicalSize {
			temp[i] = l.data[i]
		}

		l.data = temp
		l.physicalSize = newCapacity
	}

	for i := l.logicalSize; i > index; i-- {
		l.data[i] = l.data[i-1]
	}

	l.data[index] = element

	l.logicalSize++
}

func (l *ArrayList[T]) Remove(index int) {
	if index < 0 || index >= l.logicalSize {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	for i := index + 1; i < l.logicalSize; i++ {
		l.data[i-1] = l.data[i]
	}

	l.logicalSize--

	l.data[l.logicalSize] = *new(T)
}

func (l *ArrayList[T]) Front() T {
	return l.Get(0)
}

func (l *ArrayList[T]) Back() T {
	return l.Get(l.logicalSize - 1)
}

func (l *ArrayList[T]) Prepend(element T) {
	l.Insert(0, element)
}

func (l *ArrayList[T]) Append(element T) {
	l.Insert(l.logicalSize, element)
}

func (l *ArrayList[T]) Poll() {
	l.Remove(0)
}

func (l *ArrayList[T]) Eject() {
	l.Remove(l.logicalSize - 1)
}

func (l *ArrayList[T]) Capacity() int {
	return l.physicalSize
}

func (l *ArrayList[T]) Shrink() {
	temp := make([]T, l.logicalSize)

	for i := range l.logicalSize {
		temp[i] = l.data[i]
	}

	l.data = temp
	l.physicalSize = l.logicalSize
}
