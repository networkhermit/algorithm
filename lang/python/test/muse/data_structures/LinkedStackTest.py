from muse.data_structures.LinkedStack import LinkedStack
from muse.util import TestRunner

def testLinkedStack() -> bool:
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

    return stack.isEmpty()

def main() -> None:
    TestRunner.parseTest(testLinkedStack())

if __name__ == "__main__":
    main()
