from collections.abc import Callable

from muse.data_structures import List, Queue, Resizable, Stack


def list_derive[T: List](factory: type[T]) -> Callable[[], bool]:
    def f() -> bool:
        size = 8192

        list = factory()

        for i in range(1, size + 1):
            list.append(i)

        if isinstance(list, Resizable):
            list.shrink()

        if list.size() != size:
            return False

        if isinstance(list, Resizable):
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

        if isinstance(list, Resizable):
            list.shrink()

        if not list.is_empty():
            return False

        if isinstance(list, Resizable):
            return list.capacity() == 0
        else:
            return True

    return f


def queue_derive[T: Queue](factory: type[T]) -> Callable[[], bool]:
    def f() -> bool:
        size = 8192

        queue = factory()

        for i in range(1, size + 1):
            queue.enqueue(i)

        if isinstance(queue, Resizable):
            queue.shrink()

        if queue.size() != size:
            return False

        if isinstance(queue, Resizable):
            if queue.capacity() != size:
                return False

        for i in range(1, size + 1):
            if queue.peek() != i:
                return False
            queue.dequeue()

        if isinstance(queue, Resizable):
            queue.shrink()

        if not queue.is_empty():
            return False

        if isinstance(queue, Resizable):
            return queue.capacity() == 0
        else:
            return True

    return f


def stack_derive[T: Stack](factory: type[T]) -> Callable[[], bool]:
    def f() -> bool:
        size = 8192

        stack = factory()

        for i in range(1, size + 1):
            stack.push(i)

        if isinstance(stack, Resizable):
            stack.shrink()

        if stack.size() != size:
            return False

        if isinstance(stack, Resizable):
            if stack.capacity() != size:
                return False

        for i in range(size, 0, -1):
            if stack.peek() != i:
                return False
            stack.pop()

        if isinstance(stack, Resizable):
            stack.shrink()

        if not stack.is_empty():
            return False

        if isinstance(stack, Resizable):
            return stack.capacity() == 0
        else:
            return True

    return f
