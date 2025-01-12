from muse.data_structures import tests
from muse.data_structures.linked_queue import LinkedQueue
from muse.util import test_runner


def test_linked_queue() -> bool:
    return tests.queue_derive(LinkedQueue)()


def main() -> None:
    test_runner.pick(test_linked_queue)


if __name__ == "__main__":
    main()
