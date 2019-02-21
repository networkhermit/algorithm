package ArrayList

import "errors"

const DEFAULT_CAPACITY int = 64

type ArrayList struct {
    data []int
    logicalSize int
    physicalSize int
}

func New(physicalSize int) *ArrayList {
    list := &ArrayList{data: nil, logicalSize: 0, physicalSize: DEFAULT_CAPACITY}

    if physicalSize > 1 {
        list.physicalSize = physicalSize
    }
    list.data = make([]int, list.physicalSize)

    return list
}

func (list *ArrayList) Size() int {
    return list.logicalSize
}

func (list *ArrayList) IsEmpty() bool {
    return list.logicalSize == 0
}

func (list *ArrayList) Get(index int) int {
    if index < 0 || index >= list.logicalSize {
        panic(errors.New("[PANIC - IndexOutOfBounds]"))
    }

    return list.data[index]
}

func (list *ArrayList) Set(index int, element int) {
    if index < 0 || index >= list.logicalSize {
        panic(errors.New("[PANIC - IndexOutOfBounds]"))
    }

    list.data[index] = element
}

func (list *ArrayList) Insert(index int, element int) {
    if index < 0 || index > list.logicalSize {
        panic(errors.New("[PANIC - IndexOutOfBounds]"))
    }

    if list.logicalSize == list.physicalSize {
        newCapacity := DEFAULT_CAPACITY

        if list.physicalSize >= DEFAULT_CAPACITY {
            newCapacity = list.physicalSize + (list.physicalSize >> 1)
        }

        temp := make([]int, newCapacity)

        for i, length := 0, list.logicalSize; i < length; i++ {
            temp[i] = list.data[i]
        }

        list.data = temp
        list.physicalSize = newCapacity
    }

    for i := list.logicalSize; i > index; i-- {
        list.data[i] = list.data[i - 1]
    }

    list.data[index] = element

    list.logicalSize += 1
}

func (list *ArrayList) Remove(index int) {
    if index < 0 || index >= list.logicalSize {
        panic(errors.New("[PANIC - IndexOutOfBounds]"))
    }

    for i := index + 1; i < list.logicalSize; i++ {
        list.data[i - 1] = list.data[i]
    }

    list.logicalSize -= 1

    list.data[list.logicalSize] = int(0)
}

func (list *ArrayList) Front() int {
    return list.Get(0)
}

func (list *ArrayList) Back() int {
    return list.Get(list.logicalSize - 1)
}

func (list *ArrayList) Prepend(element int) {
    list.Insert(0, element)
}

func (list *ArrayList) Append(element int) {
    list.Insert(list.logicalSize, element)
}

func (list *ArrayList) Poll() {
    list.Remove(0)
}

func (list *ArrayList) Eject() {
    list.Remove(list.logicalSize - 1)
}

func (list *ArrayList) Capacity() int {
    return list.physicalSize
}

func (list *ArrayList) Shrink() {
    temp := make([]int, list.logicalSize)

    for i, length := 0, list.logicalSize; i < length; i++ {
        temp[i] = list.data[i]
    }

    list.data = temp
    list.physicalSize = list.logicalSize
}
