from muse.data_structures.array_list import ArrayList
from muse.util import test_runner


def test_array_list() -> bool:
    size = 8192

    list = ArrayList(0)

    for i in range(1, size + 1):
        list.append(i)

    list.shrink()

    if list.size() != size:
        return False

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

    list.shrink()

    if not list.is_empty():
        return False

    return list.capacity() == 0


def main() -> None:
    test_runner.pick(test_array_list)


if __name__ == "__main__":
    main()
