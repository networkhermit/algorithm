package data_structures

type Queue interface {
	Size() int

	IsEmpty() bool

	Peek() int

	Enqueue(element int)

	Dequeue()
}
