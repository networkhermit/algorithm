from muse.algorithms.sorting import selection_sort, tests
from muse.util import test_runner


def test_selection_sort() -> bool:
    sort = selection_sort.sort
    return (
        tests.derive_decreasing(sort, 10000)()
        and tests.derive_empty(sort)()
        and tests.derive_identical(sort, 10000)()
        and tests.derive_increasing(sort, 10000)()
        and tests.derive_random(sort, 10000)()
    )


def main() -> None:
    test_runner.pick(test_selection_sort)


if __name__ == "__main__":
    main()
