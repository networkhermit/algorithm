package main

import (
    "muse/data_structures/ArrayStack"
    "muse/util/TestRunner"
)

func testArrayStack() bool {
    size := 8192

    stack := ArrayStack.New(0)

    for i := 1; i <= size; i++ {
        stack.Push(i)
    }

    stack.Shrink()

    if stack.Size() != size {
        return false
    }

    if stack.Capacity() != size {
        return false
    }

    for i := size; i > 0; i-- {
        if stack.Peek() != i {
            return false
        }
        stack.Pop()
    }

    stack.Shrink()

    if !stack.IsEmpty() {
        return false
    }

    return stack.Capacity() == 0
}

func main() {
    TestRunner.ParseTest(testArrayStack())
}
