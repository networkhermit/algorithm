from muse.data_structures import tests
from muse.data_structures.array_queue import ArrayQueue
from muse.util import test_runner


def test_array_queue() -> bool:
    return tests.queue_derive(ArrayQueue)()


def main() -> None:
    test_runner.pick(test_array_queue)


if __name__ == "__main__":
    main()
