package SinglyLinkedList

import "errors"

type Node struct {
    data int
    next *Node
}

type SinglyLinkedList struct {
    head *Node
    length int
}

func New() *SinglyLinkedList {
    return &SinglyLinkedList{head: nil, length: 0}
}

func (list *SinglyLinkedList) Size() int {
    return list.length
}

func (list *SinglyLinkedList) IsEmpty() bool {
    return list.length == 0
}

func (list *SinglyLinkedList) Get(index int) int {
    if index < 0 || index >= list.length {
        panic(errors.New("[PANIC - IndexOutOfBounds]"))
    }

    cursor := list.head

    for i := 0; i < index; i++ {
        cursor = cursor.next
    }

    return cursor.data
}

func (list *SinglyLinkedList) Set(index int, element int) {
    if index < 0 || index >= list.length {
        panic(errors.New("[PANIC - IndexOutOfBounds]"))
    }

    cursor := list.head

    for i := 0; i < index; i++ {
        cursor = cursor.next
    }

    cursor.data = element
}

func (list *SinglyLinkedList) Insert(index int, element int) {
    if index < 0 || index > list.length {
        panic(errors.New("[PANIC - IndexOutOfBounds]"))
    }

    node := &Node{data: element, next: nil}

    if index == 0 {
        if list.length != 0 {
            node.next = list.head
        }
        list.head = node
    } else {
        cursor := list.head
        for i, bound := 0, index - 1; i < bound; i++ {
            cursor = cursor.next
        }
        node.next = cursor.next
        cursor.next = node
    }

    list.length++
}

func (list *SinglyLinkedList) Remove(index int) {
    if index < 0 || index >= list.length {
        panic(errors.New("[PANIC - IndexOutOfBounds]"))
    }

    var target *Node

    if index == 0 {
        target = list.head
        if list.length == 1 {
            list.head = nil
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
    }

    target.data = int(0)

    list.length--
}

func (list *SinglyLinkedList) Front() int {
    return list.Get(0)
}

func (list *SinglyLinkedList) Back() int {
    return list.Get(list.length - 1)
}

func (list *SinglyLinkedList) Prepend(element int) {
    list.Insert(0, element)
}

func (list *SinglyLinkedList) Append(element int) {
    list.Insert(list.length, element)
}

func (list *SinglyLinkedList) Poll() {
    list.Remove(0)
}

func (list *SinglyLinkedList) Eject() {
    list.Remove(list.length - 1)
}
