from muse.algorithms.sorting import insertion_sort, tests
from muse.util import test_runner


def test_insertion_sort() -> bool:
    sort = insertion_sort.sort
    return (
        tests.derive_decreasing(sort, 10000)()
        and tests.derive_empty(sort)()
        and tests.derive_identical(sort, 100000)()
        and tests.derive_increasing(sort, 100000)()
        and tests.derive_random(sort, 10000)()
    )


def main() -> None:
    test_runner.pick(test_insertion_sort)


if __name__ == "__main__":
    main()
