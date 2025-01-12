from muse.data_structures import tests
from muse.data_structures.linked_list import LinkedList
from muse.util import test_runner


def test_linked_list() -> bool:
    return tests.list_derive(LinkedList)()


def main() -> None:
    test_runner.pick(test_linked_list)


if __name__ == "__main__":
    main()
