package data_structures

type Stack interface {
	Size() int

	IsEmpty() bool

	Peek() int

	Push(element int)

	Pop()
}
