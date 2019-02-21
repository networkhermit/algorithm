package IList

type IList interface {

    Size() int

    IsEmpty() bool

    Get(index int) int

    Set(index int, element int)

    Insert(index int, element int)

    Remove(index int)

    Front() int

    Back() int

    Prepend() int

    Append() int

    Poll()

    Eject()
}
