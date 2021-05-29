package main

import (
    "vac.fun/data_structures/SinglyLinkedList"
    "vac.fun/util/TestRunner"
)

func testSinglyLinkedList() bool {
    size := 8192

    list := SinglyLinkedList.New()

    for i := 1; i <= size; i++ {
        list.Append(i)
    }

    if list.Size() != size {
        return false
    }

    for i := 0; i < size; i++ {
        if list.Get(i) != i + 1 {
            return false
        }
    }

    for i := 0; i < size; i++ {
        list.Set(i, size - i)
    }

    for i := 0; i < size; i++ {
        if list.Get(i) != size - i {
            return false
        }
    }

    var x int
    var y int

    for i, j := 0, size - 1; i < j; i, j = i + 1, j - 1 {
        x = list.Get(i)
        y = list.Get(j)

        list.Remove(i)
        list.Insert(i, y)
        list.Remove(j)
        list.Insert(j, x)
    }

    for i := 0; i < size; i++ {
        if list.Get(i) != i + 1 {
            return false
        }
    }

    for i := size; i >= 1; i-- {
        if list.Back() != i {
            return false
        }
        list.Eject()
    }

    return list.IsEmpty()
}

func main() {
    TestRunner.ParseTest(testSinglyLinkedList())
}
