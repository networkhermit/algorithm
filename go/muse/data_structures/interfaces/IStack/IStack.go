package IStack

type IStack interface {
	Size() int

	IsEmpty() bool

	Peek() int

	Push(element int)

	Pop()
}
