from muse.data_structures.array_stack import ArrayStack
from muse.util import test_runner


def test_array_stack() -> bool:
    size = 8192

    stack = ArrayStack(0)

    for i in range(1, size + 1):
        stack.push(i)

    stack.shrink()

    if stack.size() != size:
        return False

    if stack.capacity() != size:
        return False

    for i in range(size, 0, -1):
        if stack.peek() != i:
            return False
        stack.pop()

    stack.shrink()

    if not stack.is_empty():
        return False

    return stack.capacity() == 0


def main() -> None:
    test_runner.pick(test_array_stack)


if __name__ == "__main__":
    main()
