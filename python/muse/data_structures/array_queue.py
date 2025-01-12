from typing import Any


class ArrayQueue:
    DEFAULT_CAPACITY = 64

    def __init__(self, physicalSize: int = 0):
        self.data: list[Any] = []
        self.front = 0
        self.logical_size = 0
        self.physical_size = self.DEFAULT_CAPACITY

        if physicalSize > 1:
            self.physical_size = physicalSize
        self.data = [0] * self.physical_size

    def size(self) -> int:
        return self.logical_size

    def is_empty(self) -> bool:
        return self.logical_size == 0

    def peek(self) -> Any:
        if self.logical_size == 0:
            raise RuntimeError("[PANIC - NoSuchElement]")

        return self.data[self.front]

    def enqueue(self, element: Any) -> None:
        if self.logical_size == self.physical_size:
            newCapacity = self.DEFAULT_CAPACITY

            if self.physical_size >= self.DEFAULT_CAPACITY:
                newCapacity = self.physical_size + (self.physical_size >> 1)

            temp = [0] * newCapacity

            cursor = self.front

            for i in range(self.logical_size):
                if cursor == self.physical_size:
                    cursor = 0
                temp[i] = self.data[cursor]
                cursor += 1

            self.data = temp
            self.front = 0
            self.physical_size = newCapacity

        self.data[(self.front + self.logical_size) % self.physical_size] = element

        self.logical_size += 1

    def dequeue(self) -> None:
        if self.logical_size == 0:
            raise RuntimeError("[PANIC - NoSuchElement]")

        self.data[self.front] = None

        self.front = (self.front + 1) % self.physical_size

        self.logical_size -= 1

    def capacity(self) -> int:
        return self.physical_size

    def shrink(self) -> None:
        temp = [0] * self.logical_size

        cursor = self.front

        for i in range(self.logical_size):
            if cursor == self.physical_size:
                cursor = 0
            temp[i] = self.data[cursor]
            cursor += 1

        self.data = temp
        self.front = 0
        self.physical_size = self.logical_size
