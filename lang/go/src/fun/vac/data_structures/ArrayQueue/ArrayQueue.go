package ArrayQueue

import "errors"

const DEFAULT_CAPACITY int = 64

type ArrayQueue struct {
    data []int
    front int
    logicalSize int
    physicalSize int
}

func New(physicalSize int) *ArrayQueue {
    queue := &ArrayQueue{data: nil, logicalSize: 0, physicalSize: DEFAULT_CAPACITY}

    if physicalSize > 1 {
        queue.physicalSize = physicalSize
    }
    queue.data = make([]int, queue.physicalSize)

    return queue
}

func (queue *ArrayQueue) Size() int {
    return queue.logicalSize
}

func (queue *ArrayQueue) IsEmpty() bool {
    return queue.logicalSize == 0
}

func (queue *ArrayQueue) Peek() int {
    if queue.logicalSize == 0 {
        panic(errors.New("[PANIC - NoSuchElement]"))
    }

    return queue.data[queue.front]
}

func (queue *ArrayQueue) Enqueue(element int) {
    if queue.logicalSize == queue.physicalSize {
        newCapacity := DEFAULT_CAPACITY

        if queue.physicalSize >= DEFAULT_CAPACITY {
            newCapacity = queue.physicalSize + (queue.physicalSize >> 1)
        }

        temp := make([]int, newCapacity)

        cursor := queue.front

        for i, length := 0, queue.logicalSize; i < length; i++ {
            if cursor == queue.physicalSize {
                cursor = 0
            }
            temp[i] = queue.data[cursor]
            cursor += 1
        }

        queue.data = temp
        queue.front = 0
        queue.physicalSize = newCapacity
    }

    queue.data[(queue.front + queue.logicalSize) % queue.physicalSize] = element

    queue.logicalSize += 1
}

func (queue *ArrayQueue) Dequeue() {
    if queue.logicalSize == 0 {
        panic(errors.New("[PANIC - NoSuchElement]"))
    }

    queue.data[queue.front] = int(0)

    queue.front = (queue.front + 1) % queue.physicalSize

    queue.logicalSize -= 1
}

func (queue *ArrayQueue) Capacity() int {
    return queue.physicalSize
}

func (queue *ArrayQueue) Shrink() {
    temp := make([]int, queue.logicalSize)

    cursor := queue.front

    for i, length := 0, queue.logicalSize; i < length; i++ {
        if cursor == queue.physicalSize {
            cursor = 0
        }
        temp[i] = queue.data[cursor]
        cursor += 1
    }

    queue.data = temp
    queue.front = 0
    queue.physicalSize = queue.logicalSize
}
