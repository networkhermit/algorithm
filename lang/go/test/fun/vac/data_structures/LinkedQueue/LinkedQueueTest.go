package main

import (
    "fun/vac/data_structures/LinkedQueue"
    "fun/vac/util/TestRunner"
)

func testLinkedQueue() bool {
    size := 8192

    queue := LinkedQueue.New()

    for i := 1; i <= size; i++ {
        queue.Enqueue(i)
    }

    if queue.Size() != size {
        return false
    }

    for i := 1; i <= size; i++ {
        if queue.Peek() != i {
            return false
        }
        queue.Dequeue()
    }

    if !queue.IsEmpty() {
        return false
    }

    return true
}

func main() {
    TestRunner.ParseTest(testLinkedQueue())
}
