from collections.abc import Callable
from typing import Any


def list_derive(factory: type[Any]) -> Callable[[], bool]:
    def f() -> bool:
        size = 8192

        list = factory()

        for i in range(1, size + 1):
            list.append(i)

        if hasattr(list, "shrink") and callable(list.shrink):
            list.shrink()

        if list.size() != size:
            return False

        if hasattr(list, "capacity") and callable(list.capacity):
            if list.capacity() != size:
                return False

        for i in range(size):
            if list.get(i) != i + 1:
                return False

        for i in range(size):
            list.set(i, size - i)

        for i in range(size):
            if list.get(i) != size - i:
                return False

        i, j = 0, size - 1
        while i < j:
            x = list.get(i)
            y = list.get(j)

            list.remove(i)
            list.insert(i, y)
            list.remove(j)
            list.insert(j, x)
            i, j = i + 1, j - 1

        for i in range(size):
            if list.get(i) != i + 1:
                return False

        for i in range(size, 0, -1):
            if list.back() != i:
                return False
            list.eject()

        if hasattr(list, "shrink") and callable(list.shrink):
            list.shrink()

        if not list.is_empty():
            return False

        if hasattr(list, "capacity") and callable(list.capacity):
            return list.capacity() == 0
        else:
            return True

    return f


def queue_derive(factory: type[Any]) -> Callable[[], bool]:
    def f() -> bool:
        size = 8192

        queue = factory()

        for i in range(1, size + 1):
            queue.enqueue(i)

        if hasattr(queue, "shrink") and callable(queue.shrink):
            queue.shrink()

        if queue.size() != size:
            return False

        if hasattr(queue, "capacity") and callable(queue.capacity):
            if queue.capacity() != size:
                return False

        for i in range(1, size + 1):
            if queue.peek() != i:
                return False
            queue.dequeue()

        if hasattr(queue, "shrink") and callable(queue.shrink):
            queue.shrink()

        if not queue.is_empty():
            return False

        if hasattr(queue, "capacity") and callable(queue.capacity):
            return queue.capacity() == 0
        else:
            return True

    return f


def stack_derive(factory: type[Any]) -> Callable[[], bool]:
    def f() -> bool:
        size = 8192

        stack = factory()

        for i in range(1, size + 1):
            stack.push(i)

        if hasattr(stack, "shrink") and callable(stack.shrink):
            stack.shrink()

        if stack.size() != size:
            return False

        if hasattr(stack, "capacity") and callable(stack.capacity):
            if stack.capacity() != size:
                return False

        for i in range(size, 0, -1):
            if stack.peek() != i:
                return False
            stack.pop()

        if hasattr(stack, "shrink") and callable(stack.shrink):
            stack.shrink()

        if not stack.is_empty():
            return False

        if hasattr(stack, "capacity") and callable(stack.capacity):
            return stack.capacity() == 0
        else:
            return True

    return f
