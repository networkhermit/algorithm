from fun.vac.data_structures.LinkedQueue import LinkedQueue
from fun.vac.util import TestRunner

def testLinkedQueue()-> bool:
    size = 8192

    queue = LinkedQueue()

    for i in range(1, size + 1):
        queue.enqueue(i)

    if queue.size() != size:
        return False

    for i in range(1, size + 1):
        if queue.peek() != i:
            return False
        queue.dequeue()

    if not queue.isEmpty():
        return False

    return True

if __name__ == "__main__":
    TestRunner.parseTest(testLinkedQueue())
