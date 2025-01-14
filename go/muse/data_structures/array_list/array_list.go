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

func (list *ArrayList[T]) Size() int {
	return list.logicalSize
}

func (list *ArrayList[T]) IsEmpty() bool {
	return list.logicalSize == 0
}

func (list *ArrayList[T]) Get(index int) T {
	if index < 0 || index >= list.logicalSize {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	return list.data[index]
}

func (list *ArrayList[T]) Set(index int, element T) {
	if index < 0 || index >= list.logicalSize {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	list.data[index] = element
}

func (list *ArrayList[T]) Insert(index int, element T) {
	if index < 0 || index > list.logicalSize {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	if list.logicalSize == list.physicalSize {
		newCapacity := DEFAULT_CAPACITY

		if list.physicalSize >= DEFAULT_CAPACITY {
			newCapacity = list.physicalSize + (list.physicalSize >> 1)
		}

		temp := make([]T, newCapacity)

		for i := range list.logicalSize {
			temp[i] = list.data[i]
		}

		list.data = temp
		list.physicalSize = newCapacity
	}

	for i := list.logicalSize; i > index; i-- {
		list.data[i] = list.data[i-1]
	}

	list.data[index] = element

	list.logicalSize++
}

func (list *ArrayList[T]) Remove(index int) {
	if index < 0 || index >= list.logicalSize {
		panic(errors.New("[PANIC - IndexOutOfBounds]"))
	}

	for i := index + 1; i < list.logicalSize; i++ {
		list.data[i-1] = list.data[i]
	}

	list.logicalSize--

	list.data[list.logicalSize] = *new(T)
}

func (list *ArrayList[T]) Front() T {
	return list.Get(0)
}

func (list *ArrayList[T]) Back() T {
	return list.Get(list.logicalSize - 1)
}

func (list *ArrayList[T]) Prepend(element T) {
	list.Insert(0, element)
}

func (list *ArrayList[T]) Append(element T) {
	list.Insert(list.logicalSize, element)
}

func (list *ArrayList[T]) Poll() {
	list.Remove(0)
}

func (list *ArrayList[T]) Eject() {
	list.Remove(list.logicalSize - 1)
}

func (list *ArrayList[T]) Capacity() int {
	return list.physicalSize
}

func (list *ArrayList[T]) Shrink() {
	temp := make([]T, list.logicalSize)

	for i := range list.logicalSize {
		temp[i] = list.data[i]
	}

	list.data = temp
	list.physicalSize = list.logicalSize
}
