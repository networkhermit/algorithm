class ArrayStack[T]:
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

    def peek(self) -> T:
        if self.logical_size == 0:
            raise RuntimeError("[PANIC - NoSuchElement]")

        return self.data[self.logical_size - 1]

    def push(self, element: T) -> None:
        if self.logical_size == self.physical_size:
            newCapacity = self.DEFAULT_CAPACITY

            if self.physical_size >= self.DEFAULT_CAPACITY:
                newCapacity = self.physical_size + (self.physical_size >> 1)

            temp: list[T | None] = [None] * newCapacity

            for i in range(self.logical_size):
                temp[i] = self.data[i]

            self.data = temp
            self.physical_size = newCapacity

        self.data[self.logical_size] = element

        self.logical_size += 1

    def pop(self) -> None:
        if self.logical_size == 0:
            raise RuntimeError("[PANIC - NoSuchElement]")

        self.logical_size -= 1

        self.data[self.logical_size] = None

    def capacity(self) -> int:
        return self.physical_size

    def shrink(self) -> None:
        temp: list[T | None] = [None] * self.logical_size

        for i in range(self.logical_size):
            temp[i] = self.data[i]

        self.data = temp
        self.physical_size = self.logical_size
