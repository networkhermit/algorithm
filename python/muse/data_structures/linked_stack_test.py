from muse.data_structures.linked_stack import LinkedStack
from muse.util import test_runner


def test_linked_stack() -> bool:
    size = 8192

    stack = LinkedStack()

    for i in range(1, size + 1):
        stack.push(i)

    if stack.size() != size:
        return False

    for i in range(size, 0, -1):
        if stack.peek() != i:
            return False
        stack.pop()

    return stack.is_empty()


def main() -> None:
    test_runner.pick(test_linked_stack)


if __name__ == "__main__":
    main()
