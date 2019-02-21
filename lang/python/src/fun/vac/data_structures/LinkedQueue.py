class Node:

    def __init__(self, element: object):
        self.data = element
        self.next = None

class LinkedQueue:

    def __init__(self):
        self.head = None
        self.tail = None
        self.length = 0

    def size(self) -> int:
        return self.length

    def isEmpty(self) -> bool:
        return self.length == 0

    def peek(self) -> int:
        if self.length == 0:
            raise RuntimeError("[PANIC - NoSuchElement]")

        return self.head.data

    def enqueue(self, element: object) -> None:
        node = Node(element)

        if self.length == 0:
            self.head = node
        else:
            self.tail.next = node

        self.tail = node

        self.length += 1

    def dequeue(self) -> None:
        if self.length == 0:
            raise RuntimeError("[PANIC - NoSuchElement]")

        target = self.head

        if self.length == 1:
            self.head = None
            self.tail = None
        else:
            self.head = target.next

        target.data = None

        self.length -= 1
