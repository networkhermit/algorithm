class Node:
    def __init__(self, element: object):
        self.data = element
        self.next = None

class SinglyLinkedList:
    def __init__(self):
        self.head = None
        self.length = 0

    def size(self) -> int:
        return self.length

    def isEmpty(self) -> bool:
        return self.length == 0

    def get(self, index: int) -> int:
        if index < 0 or index >= self.length:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        cursor = self.head

        for _ in range(index):
            cursor = cursor.next

        return cursor.data

    def set(self, index: int, element: object) -> None:
        if index < 0 or index >= self.length:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        cursor = self.head

        for i in range(index):
            cursor = cursor.next

        cursor.data = element

    def insert(self, index: int, element: object) -> None:
        if index < 0 or index > self.length:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        node = Node(element)

        if index == 0:
            if self.length != 0:
                node.next = self.head
            self.head = node
        else:
            cursor = self.head
            for i in range(index - 1):
                cursor = cursor.next
            node.next = cursor.next
            cursor.next = node

        self.length += 1

    def remove(self, index: int) -> None:
        if index < 0 or index >= self.length:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        target = None

        if index == 0:
            target = self.head
            if self.length == 1:
                self.head = None
            else:
                self.head = target.next
        else:
            cursor = self.head
            for i in range(index - 1):
                cursor = cursor.next
            target = cursor.next
            cursor.next = target.next

        target.data = None

        self.length -= 1

    def front(self) -> int:
        return self.get(0)

    def back(self) -> int:
        return self.get(self.length - 1)

    def prepend(self, element: object) -> None:
        self.insert(0, element)

    def append(self, element: object) -> None:
        self.insert(self.length, element)

    def poll(self) -> None:
        self.remove(0)

    def eject(self) -> None:
        self.remove(self.length - 1)
