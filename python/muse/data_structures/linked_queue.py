class Node[T]:
    def __init__(self, element: T):
        self.data: T | None = element
        self.next: Node[T] | None = None


class LinkedQueue[T]:
    def __init__(self):
        self.head = None
        self.tail = None
        self.length = 0

    def size(self) -> int:
        return self.length

    def is_empty(self) -> bool:
        return self.length == 0

    def peek(self) -> T:
        if self.length == 0:
            raise RuntimeError("[PANIC - NoSuchElement]")

        return self.head.data

    def enqueue(self, element: T) -> None:
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
