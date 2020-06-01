class ArrayList:
    DEFAULT_CAPACITY = 64

    def __init__(self, physicalSize: int = 0):
        self.data = None
        self.logicalSize = 0
        self.physicalSize = self.DEFAULT_CAPACITY

        if physicalSize > 1:
            self.physicalSize = physicalSize
        self.data = [0] * self.physicalSize

    def size(self) -> int:
        return self.logicalSize

    def isEmpty(self) -> bool:
        return self.logicalSize == 0

    def get(self, index: int) -> int:
        if index < 0 or index >= self.logicalSize:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        return self.data[index]

    def set(self, index: int, element: object) -> None:
        if index < 0 or index >= self.logicalSize:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        self.data[index] = element

    def insert(self, index: int, element: object) -> None:
        if index < 0 or index > self.logicalSize:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        if self.logicalSize == self.physicalSize:
            newCapacity = self.DEFAULT_CAPACITY

            if self.physicalSize >= self.DEFAULT_CAPACITY:
                newCapacity = self.physicalSize + (self.physicalSize >> 1)

            temp = [0] * newCapacity

            for i in range(self.logicalSize):
                temp[i] = self.data[i]

            self.data = temp
            self.physicalSize = newCapacity

        for i in range(self.logicalSize, index, -1):
            self.data[i] = self.data[i - 1]

        self.data[index] = element

        self.logicalSize += 1

    def remove(self, index: int) -> None:
        if index < 0 or index >= self.logicalSize:
            raise RuntimeError("[PANIC - IndexOutOfBounds]")

        for i in range(index + 1, self.logicalSize):
            self.data[i - 1] = self.data[i]

        self.logicalSize -= 1

        self.data[self.logicalSize] = None

    def front(self) -> int:
        return self.get(0)

    def back(self) -> int:
        return self.get(self.logicalSize - 1)

    def prepend(self, element: object) -> None:
        self.insert(0, element)

    def append(self, element: object) -> None:
        self.insert(self.logicalSize, element)

    def poll(self) -> None:
        self.remove(0)

    def eject(self) -> None:
        self.remove(self.logicalSize - 1)

    def capacity(self) -> int:
        return self.physicalSize

    def shrink(self) -> None:
        temp = [0] * self.logicalSize

        for i in range(self.logicalSize):
            temp[i] = self.data[i]

        self.data = temp
        self.physicalSize = self.logicalSize
