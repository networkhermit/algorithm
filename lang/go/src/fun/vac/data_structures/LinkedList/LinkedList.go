package LinkedList

import "errors"

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
    tail *Node
    length int
}

func New() *LinkedList {
    return &LinkedList{head: nil, tail: nil, length: 0}
}

func (list *LinkedList) Size() int {
    return list.length
}

func (list *LinkedList) IsEmpty() bool {
    return list.length == 0
}

func (list *LinkedList) Get(index int) int {
    if index < 0 || index >= list.length {
        panic(errors.New("[PANIC - IndexOutOfBounds]"))
    }

    var cursor *Node

    if index == list.length - 1 {
        cursor = list.tail
    } else {
        cursor = list.head
        for i := 0; i < index; i++ {
            cursor = cursor.next
        }
    }

    return cursor.data
}

func (list *LinkedList) Set(index int, element int) {
    if index < 0 || index >= list.length {
        panic(errors.New("[PANIC - IndexOutOfBounds]"))
    }

    var cursor *Node

    if index == list.length - 1 {
        cursor = list.tail
    } else {
        cursor = list.head
        for i := 0; i < index; i++ {
            cursor = cursor.next
        }
    }

    cursor.data = element
}

func (list *LinkedList) Insert(index int, element int) {
    if index < 0 || index > list.length {
        panic(errors.New("[PANIC - IndexOutOfBounds]"))
    }

    node := &Node{data: element, next: nil}

    if index == 0 {
        if list.length != 0 {
            node.next = list.head
        } else {
            list.tail = node
        }
        list.head = node
    } else if index == list.length {
        list.tail.next = node
        list.tail = node
    } else {
        cursor := list.head
        for i, bound := 0, index - 1; i < bound; i++ {
            cursor = cursor.next
        }
        node.next = cursor.next
        cursor.next = node
    }

    list.length += 1
}

func (list *LinkedList) Remove(index int) {
    if index < 0 || index >= list.length {
        panic(errors.New("[PANIC - IndexOutOfBounds]"))
    }

    var target *Node

    if index == 0 {
        target = list.head
        if list.length == 1 {
            list.head = nil
            list.tail = nil
        } else {
            list.head = target.next
        }
    } else {
        cursor := list.head
        for i, bound := 0, index - 1; i < bound; i++ {
            cursor = cursor.next
        }
        target = cursor.next
        cursor.next = target.next
        if index == list.length - 1 {
            list.tail = cursor
        }
    }

    target.data = int(0)

    list.length -= 1
}

func (list *LinkedList) Front() int {
    return list.Get(0)
}

func (list *LinkedList) Back() int {
    return list.Get(list.length - 1)
}

func (list *LinkedList) Prepend(element int) {
    list.Insert(0, element)
}

func (list *LinkedList) Append(element int) {
    list.Insert(list.length, element)
}

func (list *LinkedList) Poll() {
    list.Remove(0)
}

func (list *LinkedList) Eject() {
    list.Remove(list.length - 1)
}
