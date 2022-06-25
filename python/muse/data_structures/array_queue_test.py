from muse.data_structures.array_queue import ArrayQueue
from muse.util import test_runner


def test_array_queue() -> bool:
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

    if not queue.is_empty():
        return False

    return queue.capacity() == 0


def main() -> None:
    test_runner.pick(test_array_queue)


if __name__ == "__main__":
    main()
