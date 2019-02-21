class ArrayQueue:

    DEFAULT_CAPACITY = 64

    def __init__(self, physicalSize: int = 0):
        self.data = None
        self.front = 0
        self.logicalSize = 0
        self.physicalSize = self.DEFAULT_CAPACITY

        if physicalSize > 1:
            self.physicalSize = physicalSize
        self.data = [0] * self.physicalSize

    def size(self) -> int:
        return self.logicalSize

    def isEmpty(self) -> bool:
        return self.logicalSize == 0

    def peek(self) -> int:
        if self.logicalSize == 0:
            raise RuntimeError("[PANIC - NoSuchElement]")

        return self.data[self.front]

    def enqueue(self, element: object) -> None:
        if self.logicalSize == self.physicalSize:
            newCapacity = self.DEFAULT_CAPACITY

            if self.physicalSize >= self.DEFAULT_CAPACITY:
                newCapacity = self.physicalSize + (self.physicalSize >> 1)

            temp = [0] * newCapacity

            cursor = self.front

            for i in range(self.logicalSize):
                if cursor == self.physicalSize:
                    cursor = 0
                temp[i] = self.data[cursor]
                cursor += 1

            self.data = temp
            self.front = 0
            self.physicalSize = newCapacity

        self.data[(self.front + self.logicalSize) % self.physicalSize] = element

        self.logicalSize += 1

    def dequeue(self) -> None:
        if self.logicalSize == 0:
            raise RuntimeError("[PANIC - NoSuchElement]")

        self.data[self.front] = None

        self.front = (self.front + 1) % self.physicalSize

        self.logicalSize -= 1

    def capacity(self) -> int:
        return self.physicalSize

    def shrink(self) -> None:
        temp = [0] * self.logicalSize

        cursor = self.front

        for i in range(self.logicalSize):
            if cursor == self.physicalSize:
                cursor = 0
            temp[i] = self.data[cursor]
            cursor += 1

        self.data = temp
        self.front = 0
        self.physicalSize = self.logicalSize
