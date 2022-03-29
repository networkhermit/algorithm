package IQueue

type IQueue interface {
	Size() int

	IsEmpty() bool

	Peek() int

	Enqueue(element int)

	Dequeue()
}
