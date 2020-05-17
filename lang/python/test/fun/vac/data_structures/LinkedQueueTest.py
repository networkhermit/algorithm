from fun.vac.data_structures.LinkedQueue import LinkedQueue
from fun.vac.util import TestRunner

def testLinkedQueue() -> bool:
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

    return queue.isEmpty()

def main() -> None:
    TestRunner.parseTest(testLinkedQueue())

if __name__ == "__main__":
    main()
