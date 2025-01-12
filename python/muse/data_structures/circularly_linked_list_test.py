from muse.data_structures import tests
from muse.data_structures.circularly_linked_list import CircularlyLinkedList
from muse.util import test_runner


def test_circularly_linked_list() -> bool:
    return tests.list_derive(CircularlyLinkedList)()


def main() -> None:
    test_runner.pick(test_circularly_linked_list)


if __name__ == "__main__":
    main()
