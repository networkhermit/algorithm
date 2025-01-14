package data_structures

type Queue[T any] interface {
	Size() int

	IsEmpty() bool

	Peek() T

	Enqueue(element T)

	Dequeue()
}
