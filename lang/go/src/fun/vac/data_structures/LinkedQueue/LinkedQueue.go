package LinkedQueue

import "errors"

type Node struct {
    data int
    next *Node
}

type LinkedQueue struct {
    head *Node
    tail *Node
    length int
}

func New() *LinkedQueue {
    return &LinkedQueue{head: nil, tail: nil, length: 0}
}

func (queue *LinkedQueue) Size() int {
    return queue.length
}

func (queue *LinkedQueue) IsEmpty() bool {
    return queue.length == 0
}

func (queue *LinkedQueue) Peek() int {
    if queue.length == 0 {
        panic(errors.New("[PANIC - NoSuchElement]"))
    }

    return queue.head.data
}

func (queue *LinkedQueue) Enqueue(element int) {
    node := &Node{data: element, next: nil}

    if queue.length == 0 {
        queue.head = node
    } else {
        queue.tail.next = node
    }

    queue.tail = node

    queue.length += 1
}

func (queue *LinkedQueue) Dequeue() {
    if queue.length == 0 {
        panic(errors.New("[PANIC - NoSuchElement]"))
    }

    target := queue.head

    if queue.length == 1 {
        queue.head = nil
        queue.tail = nil
    } else {
        queue.head = target.next
    }

    target.data = int(0)

    queue.length -= 1
}
