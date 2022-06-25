from muse.data_structures.linked_queue import LinkedQueue
from muse.util import test_runner


def test_linked_queue() -> bool:
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

    return queue.is_empty()


def main() -> None:
    test_runner.pick(test_linked_queue)


if __name__ == "__main__":
    main()
