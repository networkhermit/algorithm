from muse.algorithms.search import binary_search
from muse.util import sequence_builder, test_runner


def test_binary_search() -> bool:
    size = 32768

    arr = [0] * size
    sequence_builder.pack_increasing(arr)

    if binary_search.find(arr, -1) != size:
        return False

    if binary_search.find(arr, 2_147_483_647) != size:
        return False

    for i, v in enumerate(arr):
        if binary_search.find(arr, v) != i:
            return False

    return True


def main() -> None:
    test_runner.pick(test_binary_search)


if __name__ == "__main__":
    main()
