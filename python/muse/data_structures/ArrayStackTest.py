from muse.data_structures.ArrayStack import ArrayStack
from muse.util import TestRunner

def testArrayStack() -> bool:
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

    if not stack.isEmpty():
        return False

    return stack.capacity() == 0

def main() -> None:
    TestRunner.pick(testArrayStack)

if __name__ == "__main__":
    main()
