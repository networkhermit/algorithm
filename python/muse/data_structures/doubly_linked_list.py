class Node:
    def __init__(self, element: object):
        self.data = element
        self.next = None
        self.prev = None


class DoublyLinkedList:
    def __init__(self):
        self.head = None
        self.tail = None
        self.length = 0

    def size(self) -> int:
        return self.length

    def is_empty(self) -> bool:
        return self.length == 0

    def get(self, index: int) -> int:
        if index < 0 or index >= self.length:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        cursor = None

        if index < self.length >> 1:
            cursor = self.head
            for _ in range(index):
                cursor = cursor.next
        else:
            cursor = self.tail
            for _ in range(self.length - 1, index, -1):
                cursor = cursor.prev

        return cursor.data

    def set(self, index: int, element: object) -> None:
        if index < 0 or index >= self.length:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        cursor = None

        if index < self.length >> 1:
            cursor = self.head
            for i in range(index):
                cursor = cursor.next
        else:
            cursor = self.tail
            for i in range(self.length - 1, index, -1):
                cursor = cursor.prev

        cursor.data = element

    def insert(self, index: int, element: object) -> None:
        if index < 0 or index > self.length:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        node = Node(element)

        if index == 0:
            if self.length != 0:
                node.next = self.head
                self.head.prev = node
            else:
                self.tail = node
            self.head = node
        elif index == self.length:
            node.prev = self.tail
            self.tail.next = node
            self.tail = node
        else:
            cursor = None
            if index < self.length >> 1:
                cursor = self.head
                for i in range(index):
                    cursor = cursor.next
            else:
                cursor = self.tail
                for i in range(self.length - 1, index, -1):
                    cursor = cursor.prev
            node.next = cursor
            node.prev = cursor.prev
            cursor.prev.next = node
            cursor.prev = node

        self.length += 1

    def remove(self, index: int) -> None:
        if index < 0 or index >= self.length:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        target = None

        if index == 0:
            target = self.head
            if self.length == 1:
                self.head = None
                self.tail = None
            else:
                target.next.prev = None
                self.head = target.next
        elif index == self.length - 1:
            target = self.tail
            target.prev.next = None
            self.tail = target.prev
        else:
            if index < self.length >> 1:
                target = self.head
                for i in range(index):
                    target = target.next
            else:
                target = self.tail
                for i in range(self.length - 1, index, -1):
                    target = target.prev
            target.prev.next = target.next
            target.next.prev = target.prev

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
