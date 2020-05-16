package main

import (
    "fun/vac/data_structures/ArrayList"
    "fun/vac/util/TestRunner"
)

func testArrayList() bool {
    size := 8192

    list := ArrayList.New(0)

    for i := 1; i <= size; i++ {
        list.Append(i)
    }

    list.Shrink()

    if list.Size() != size {
        return false
    }

    if list.Capacity() != size {
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

    list.Shrink()

    if !list.IsEmpty() {
        return false
    }

    return list.Capacity() == 0
}

func main() {
    TestRunner.ParseTest(testArrayList())
}
