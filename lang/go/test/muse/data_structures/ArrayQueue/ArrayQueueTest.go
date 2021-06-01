package main

import (
    "muse/data_structures/ArrayQueue"
    "muse/util/TestRunner"
)

func testArrayQueue() bool {
    size := 8192

    queue := ArrayQueue.New(0)

    for i := 1; i <= size; i++ {
        queue.Enqueue(i)
    }

    queue.Shrink()

    if queue.Size() != size {
        return false
    }

    if queue.Capacity() != size {
        return false
    }

    for i := 1; i <= size; i++ {
        if queue.Peek() != i {
            return false
        }
        queue.Dequeue()
    }

    queue.Shrink()

    if !queue.IsEmpty() {
        return false
    }

    return queue.Capacity() == 0
}

func main() {
    TestRunner.ParseTest(testArrayQueue())
}
