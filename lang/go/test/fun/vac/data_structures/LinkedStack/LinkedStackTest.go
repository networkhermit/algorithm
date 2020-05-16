package main

import (
    "fun/vac/data_structures/LinkedStack"
    "fun/vac/util/TestRunner"
)

func testLinkedStack() bool {
    size := 8192

    stack := LinkedStack.New()

    for i := 1; i <= size; i++ {
        stack.Push(i)
    }

    if stack.Size() != size {
        return false
    }

    for i := size; i > 0; i-- {
        if stack.Peek() != i {
            return false
        }
        stack.Pop()
    }

    return stack.IsEmpty()
}

func main() {
    TestRunner.ParseTest(testLinkedStack())
}
