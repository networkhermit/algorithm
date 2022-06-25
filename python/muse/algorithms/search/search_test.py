from muse.algorithms import search
from muse.util import sequence_builder, test_runner


def test_binary_search() -> bool:
    size = 32768

    arr = [0] * size
    sequence_builder.pack_increasing(arr)

    if search.binarysearch(arr, -1) != size:
        return False

    if search.binarysearch(arr, 2_147_483_647) != size:
        return False

    for i, v in enumerate(arr):
        if search.binarysearch(arr, v) != i:
            return False

    return True


def test_linear_search() -> bool:
    size = 32768

    arr = [0] * size
    sequence_builder.pack_increasing(arr)

    if search.linearsearch(arr, -1) != size:
        return False

    if search.linearsearch(arr, 2_147_483_647) != size:
        return False

    for i, v in enumerate(arr):
        if search.linearsearch(arr, v) != i:
            return False

    return True


def main() -> None:
    test_runner.pick(test_binary_search)

    test_runner.pick(test_linear_search)


if __name__ == "__main__":
    main()
