package LinkedStack

import "errors"

type Node struct {
    data int
    next *Node
}

type LinkedStack struct {
    head *Node
    tail *Node
    length int
}

func New() *LinkedStack {
    return &LinkedStack{head: nil, tail: nil, length: 0}
}

func (stack *LinkedStack) Size() int {
    return stack.length
}

func (stack *LinkedStack) IsEmpty() bool {
    return stack.length == 0
}

func (stack *LinkedStack) Peek() int {
    if stack.length == 0 {
        panic(errors.New("[PANIC - NoSuchElement]"))
    }

    return stack.tail.data
}

func (stack *LinkedStack) Push(element int) {
    node := &Node{data: element, next: nil}

    if stack.length == 0 {
        stack.head = node
    } else {
        stack.tail.next = node
    }

    stack.tail = node

    stack.length++
}

func (stack *LinkedStack) Pop() {
    if stack.length == 0 {
        panic(errors.New("[PANIC - NoSuchElement]"))
    }

    target := stack.tail

    if stack.length == 1 {
        stack.head = nil
        stack.tail = nil
    } else {
        cursor := stack.head
        for i, bound := 0, stack.length - 2; i < bound; i++ {
            cursor = cursor.next
        }
        cursor.next = nil
        stack.tail = cursor
    }

    target.data = int(0)

    stack.length--
}
