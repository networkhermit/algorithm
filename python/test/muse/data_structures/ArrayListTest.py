from muse.data_structures.ArrayList import ArrayList
from muse.util import TestRunner

def testArrayList() -> bool:
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

    x = 0
    y = 0

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

    if not list.isEmpty():
        return False

    return list.capacity() == 0

def main() -> None:
    TestRunner.parseTest(testArrayList())

if __name__ == "__main__":
    main()
