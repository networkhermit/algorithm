package data_structures

type List[T any] interface {
	Size() int

	IsEmpty() bool

	Get(index int) T

	Set(index int, element T)

	Insert(index int, element T)

	Remove(index int)

	Front() T

	Back() T

	Prepend(element T)

	Append(element T)

	Poll()

	Eject()
}
