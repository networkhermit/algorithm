package main

import (
    "vac.fun/data_structures/LinkedQueue"
    "vac.fun/util/TestRunner"
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

    return queue.IsEmpty()
}

func main() {
    TestRunner.ParseTest(testLinkedQueue())
}
