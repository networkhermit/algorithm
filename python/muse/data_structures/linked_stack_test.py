from muse.data_structures import tests
from muse.data_structures.linked_stack import LinkedStack
from muse.util import test_runner


def test_linked_stack() -> bool:
    return tests.stack_derive(LinkedStack)()


def main() -> None:
    test_runner.pick(test_linked_stack)


if __name__ == "__main__":
    main()
