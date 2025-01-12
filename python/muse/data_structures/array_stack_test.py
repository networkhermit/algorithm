from muse.data_structures import tests
from muse.data_structures.array_stack import ArrayStack
from muse.util import test_runner


def test_array_stack() -> bool:
    return tests.stack_derive(ArrayStack)()


def main() -> None:
    test_runner.pick(test_array_stack)


if __name__ == "__main__":
    main()
