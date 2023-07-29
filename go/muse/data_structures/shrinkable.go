package data_structures

type Shrinkable interface {
	Capacity() int
	Shrink()
}
