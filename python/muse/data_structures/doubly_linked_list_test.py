from muse.data_structures import tests
from muse.data_structures.doubly_linked_list import DoublyLinkedList
from muse.util import test_runner


def test_doubly_linked_list() -> bool:
    return tests.list_derive(DoublyLinkedList)()


def main() -> None:
    test_runner.pick(test_doubly_linked_list)


if __name__ == "__main__":
    main()
