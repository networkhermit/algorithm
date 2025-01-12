from typing import Any


class ArrayList:
    DEFAULT_CAPACITY = 64

    def __init__(self, physicalSize: int = 0):
        self.data: list[Any] = []
        self.logical_size = 0
        self.physical_size = self.DEFAULT_CAPACITY

        if physicalSize > 1:
            self.physical_size = physicalSize
        self.data = [0] * self.physical_size

    def size(self) -> int:
        return self.logical_size

    def is_empty(self) -> bool:
        return self.logical_size == 0

    def get(self, index: int) -> Any:
        if index < 0 or index >= self.logical_size:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        return self.data[index]

    def set(self, index: int, element: Any) -> None:
        if index < 0 or index >= self.logical_size:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        self.data[index] = element

    def insert(self, index: int, element: Any) -> None:
        if index < 0 or index > self.logical_size:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        if self.logical_size == self.physical_size:
            newCapacity = self.DEFAULT_CAPACITY

            if self.physical_size >= self.DEFAULT_CAPACITY:
                newCapacity = self.physical_size + (self.physical_size >> 1)

            temp = [0] * newCapacity

            for i in range(self.logical_size):
                temp[i] = self.data[i]

            self.data = temp
            self.physical_size = newCapacity

        for i in range(self.logical_size, index, -1):
            self.data[i] = self.data[i - 1]

        self.data[index] = element

        self.logical_size += 1

    def remove(self, index: int) -> None:
        if index < 0 or index >= self.logical_size:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        for i in range(index + 1, self.logical_size):
            self.data[i - 1] = self.data[i]

        self.logical_size -= 1

        self.data[self.logical_size] = None

    def front(self) -> Any:
        return self.get(0)

    def back(self) -> Any:
        return self.get(self.logical_size - 1)

    def prepend(self, element: Any) -> None:
        self.insert(0, element)

    def append(self, element: Any) -> None:
        self.insert(self.logical_size, element)

    def poll(self) -> None:
        self.remove(0)

    def eject(self) -> None:
        self.remove(self.logical_size - 1)

    def capacity(self) -> int:
        return self.physical_size

    def shrink(self) -> None:
        temp = [0] * self.logical_size

        for i in range(self.logical_size):
            temp[i] = self.data[i]

        self.data = temp
        self.physical_size = self.logical_size
