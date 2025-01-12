from typing import Any


class Node:
    def __init__(self, element: Any):
        self.data = element
        self.next: Node | None = None


class LinkedStack:
    def __init__(self):
        self.head = None
        self.tail = None
        self.length = 0

    def size(self) -> int:
        return self.length

    def is_empty(self) -> bool:
        return self.length == 0

    def peek(self) -> Any:
        if self.length == 0:
            raise RuntimeError("[PANIC - NoSuchElement]")

        return self.tail.data

    def push(self, element: Any) -> None:
        node = Node(element)

        if self.length == 0:
            self.head = node
        else:
            self.tail.next = node

        self.tail = node

        self.length += 1

    def pop(self) -> None:
        if self.length == 0:
            raise RuntimeError("[PANIC - NoSuchElement]")

        target = self.tail

        if self.length == 1:
            self.head = None
            self.tail = None
        else:
            cursor = self.head
            for _ in range(self.length - 2):
                cursor = cursor.next
            cursor.next = None
            self.tail = cursor

        target.data = None

        self.length -= 1
