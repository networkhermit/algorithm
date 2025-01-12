from muse.algorithms import sorting
from muse.algorithms.sorting import tests
from muse.util import test_runner


def test_bubble_sort() -> bool:
    return tests.derive_random(sorting.bubblesort, 10000)()


def test_insertion_sort() -> bool:
    return tests.derive_random(sorting.insertionsort, 10000)()


def test_merge_sort() -> bool:
    return tests.derive_random(sorting.mergesort, 100000)()


def test_quick_sort() -> bool:
    return tests.derive_random(sorting.quicksort, 100000)()


def test_selection_sort() -> bool:
    return tests.derive_random(sorting.selectionsort, 10000)()


def main() -> None:
    test_runner.pick(test_bubble_sort)

    test_runner.pick(test_insertion_sort)

    test_runner.pick(test_merge_sort)

    test_runner.pick(test_quick_sort)

    test_runner.pick(test_selection_sort)


if __name__ == "__main__":
    main()
