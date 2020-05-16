from fun.vac.data_structures.ArrayQueue import ArrayQueue
from fun.vac.util import TestRunner

def testArrayQueue()-> bool:
    size = 8192

    queue = ArrayQueue(0)

    for i in range(1, size + 1):
        queue.enqueue(i)

    queue.shrink()

    if queue.size() != size:
        return False

    if queue.capacity() != size:
        return False

    for i in range(1, size + 1):
        if queue.peek() != i:
            return False
        queue.dequeue()

    queue.shrink()

    if not queue.isEmpty():
        return False

    return queue.capacity() == 0

if __name__ == "__main__":
    TestRunner.parseTest(testArrayQueue())
