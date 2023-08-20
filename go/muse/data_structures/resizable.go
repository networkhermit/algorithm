package data_structures

type Resizable interface {
	Capacity() int
	Shrink()
}
