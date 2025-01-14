package data_structures

type Stack[T any] interface {
	Size() int

	IsEmpty() bool

	Peek() T

	Push(element T)

	Pop()
}
