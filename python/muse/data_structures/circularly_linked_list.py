class Node[T]:
    def __init__(self, element: T):
        self.data: T | None = element
        self.next: Node[T] | None = None


class CircularlyLinkedList[T]:
    def __init__(self):
        self.tail = None
        self.length = 0

    def size(self) -> int:
        return self.length

    def is_empty(self) -> bool:
        return self.length == 0

    def get(self, index: int) -> T:
        if index < 0 or index >= self.length:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        cursor = self.tail

        if index != self.length - 1:
            for _ in range(index + 1):
                cursor = cursor.next

        return cursor.data

    def set(self, index: int, element: T) -> None:
        if index < 0 or index >= self.length:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        cursor = self.tail

        if index != self.length - 1:
            for _ in range(index + 1):
                cursor = cursor.next

        cursor.data = element

    def insert(self, index: int, element: T) -> None:
        if index < 0 or index > self.length:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        node = Node(element)

        if index == 0:
            if self.length == 0:
                node.next = node
                self.tail = node
            else:
                node.next = self.tail.next
                self.tail.next = node
        elif index == self.length:
            node.next = self.tail.next
            self.tail.next = node
            self.tail = node
        else:
            cursor = self.tail
            for _ in range(index):
                cursor = cursor.next
            node.next = cursor.next
            cursor.next = node

        self.length += 1

    def remove(self, index: int) -> None:
        if index < 0 or index >= self.length:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        target = None

        if index == 0:
            target = self.tail.next
            if self.length == 1:
                self.tail = None
            else:
                self.tail.next = target.next
        else:
            cursor = self.tail
            for _ in range(index):
                cursor = cursor.next
            target = cursor.next
            cursor.next = target.next
            if index == self.length - 1:
                self.tail = cursor

        target.data = None

        self.length -= 1

    def front(self) -> T:
        return self.get(0)

    def back(self) -> T:
        return self.get(self.length - 1)

    def prepend(self, element: T) -> None:
        self.insert(0, element)

    def append(self, element: T) -> None:
        self.insert(self.length, element)

    def poll(self) -> None:
        self.remove(0)

    def eject(self) -> None:
        self.remove(self.length - 1)
