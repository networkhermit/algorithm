class ArrayList[T]:
    DEFAULT_CAPACITY = 64

    def __init__(self, physicalSize: int = 0):
        self.data: list[T | None] = []
        self.logical_size = 0
        self.physical_size = self.DEFAULT_CAPACITY

        if physicalSize > 1:
            self.physical_size = physicalSize
        self.data = [None] * self.physical_size

    def size(self) -> int:
        return self.logical_size

    def is_empty(self) -> bool:
        return self.logical_size == 0

    def get(self, index: int) -> T:
        if index < 0 or index >= self.logical_size:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        return self.data[index]

    def set(self, index: int, element: T) -> None:
        if index < 0 or index >= self.logical_size:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        self.data[index] = element

    def insert(self, index: int, element: T) -> None:
        if index < 0 or index > self.logical_size:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        if self.logical_size == self.physical_size:
            newCapacity = self.DEFAULT_CAPACITY

            if self.physical_size >= self.DEFAULT_CAPACITY:
                newCapacity = self.physical_size + (self.physical_size >> 1)

            temp: list[T | None] = [None] * newCapacity

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

    def front(self) -> T:
        return self.get(0)

    def back(self) -> T:
        return self.get(self.logical_size - 1)

    def prepend(self, element: T) -> None:
        self.insert(0, element)

    def append(self, element: T) -> None:
        self.insert(self.logical_size, element)

    def poll(self) -> None:
        self.remove(0)

    def eject(self) -> None:
        self.remove(self.logical_size - 1)

    def capacity(self) -> int:
        return self.physical_size

    def shrink(self) -> None:
        temp: list[T | None] = [None] * self.logical_size

        for i in range(self.logical_size):
            temp[i] = self.data[i]

        self.data = temp
        self.physical_size = self.logical_size
