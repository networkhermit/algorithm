from muse.data_structures import tests
from muse.data_structures.singly_linked_list import SinglyLinkedList
from muse.util import test_runner


def test_singly_linked_list() -> bool:
    return tests.list_derive(SinglyLinkedList)()


def main() -> None:
    test_runner.pick(test_singly_linked_list)


if __name__ == "__main__":
    main()
